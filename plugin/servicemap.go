// NOTE this file was taken from https://github.com/gorilla/rpc/blob/master/map.go
// modified highly

// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugin

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

var (
	// Precompute the reflect.Type of some types
	typeOfError   = reflect.TypeOf((*error)(nil)).Elem()
	typeOfContext = reflect.TypeOf((*Context)(nil)).Elem()
)

// ----------------------------------------------------------------------------
// service
// ----------------------------------------------------------------------------

type service struct {
	name     string                    // name of service
	rcvr     reflect.Value             // receiver of methods for the service
	rcvrType reflect.Type              // type of the receiver
	methods  map[string]*serviceMethod // registered methods
}

type serviceMethod struct {
	method     reflect.Method // receiver method
	argsType   reflect.Type   // type of the request argument
	resultType reflect.Type   // type of the response argument
	readOnly   bool           // should this method be called from Contract.StaticCall?
}

// ----------------------------------------------------------------------------
// serviceMap
// ----------------------------------------------------------------------------

// serviceMap is a registry for services.
type serviceMap struct {
	mutex    sync.Mutex
	services map[string]*service
}

// register adds a new service using reflection to extract its methods.
func (m *serviceMap) Register(rcvr interface{}, name string) error {
	// Setup service.
	s := &service{
		name:     name,
		rcvr:     reflect.ValueOf(rcvr),
		rcvrType: reflect.TypeOf(rcvr),
		methods:  make(map[string]*serviceMethod),
	}
	if name == "" {
		s.name = reflect.Indirect(s.rcvr).Type().Name()
		if !isExported(s.name) {
			return fmt.Errorf("type %q is not exported", s.name)
		}
	}
	if s.name == "" {
		return fmt.Errorf("no service name for type %q",
			s.rcvrType.String())
	}
	// Setup methods.
	for i := 0; i < s.rcvrType.NumMethod(); i++ {
		method := s.rcvrType.Method(i)
		mtype := method.Type

		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}

		// Method needs four ins: receiver, plugin.Context, *args.
		if mtype.NumIn() != 3 {
			continue
		}

		// Second argument must be a pointer and must be something that implements the
		// plugin.Context interface
		contextType := mtype.In(1)
		if !contextType.Implements(typeOfContext) {
			continue
		}

		// Third argument must be a pointer and must be exported.
		args := mtype.In(2)
		if args.Kind() != reflect.Ptr || !isExportedOrBuiltin(args) {
			continue
		}

		// Method must have one or two output, if there's only one it must be error, otherwise
		// the first output must be a pointer and the second an error.
		if mtype.NumOut() != 1 && mtype.NumOut() != 2 {
			continue
		}
		if returnType := mtype.Out(mtype.NumOut() - 1); returnType != typeOfError {
			continue
		}
		if mtype.NumOut() == 2 {
			if returnType := mtype.Out(0); returnType.Kind() != reflect.Ptr || !isExportedOrBuiltin(returnType) {
				continue
			}
		}
		srvMethod := &serviceMethod{
			method:   method,
			argsType: args.Elem(),
		}
		if mtype.NumOut() == 2 {
			srvMethod.resultType = mtype.Out(0).Elem()
			// Currently tx handling methods don't return anything other than errors, so we can
			// distinguish tx handlers & query handlers based on the number of return values...
			// That may need to change if tx handlers are allowed to return non-error values at some
			// point in the future.
			srvMethod.readOnly = true
		}
		s.methods[method.Name] = srvMethod
	}
	if len(s.methods) == 0 {
		return fmt.Errorf("%q has no exported methods of suitable type",
			s.name)
	}
	// Add to the map.
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.services == nil {
		m.services = make(map[string]*service)
	} else if _, ok := m.services[s.name]; ok {
		return fmt.Errorf("service already defined: %q", s.name)
	}
	m.services[s.name] = s
	return nil
}

// Get returns a contract method matching the given name.
// The method name should be prefixed by the contract plugin name "MyContract.Method".
// If readOnly is true only methods that can be called via Contract.StaticCall will be matched,
// these methods cannot modify app state. On the other hand if readOnly is false then only
// methods that can be called via Contract.Call will be matched, these methods can modify app state.
func (m *serviceMap) Get(method string, readOnly bool) (*service, *serviceMethod, error) {
	parts := strings.Split(method, ".")
	if len(parts) != 2 {
		err := fmt.Errorf("service/method request ill-formed: %q", method)
		return nil, nil, err
	}
	m.mutex.Lock()
	service := m.services[parts[0]]
	m.mutex.Unlock()
	if service == nil {
		err := fmt.Errorf("can't find service %q", method)
		return nil, nil, err
	}
	serviceMethod := service.methods[parts[1]]
	if serviceMethod == nil {
		err := fmt.Errorf("can't find method %q", method)
		return nil, nil, err
	}
	if serviceMethod.readOnly != readOnly {
		if serviceMethod.readOnly {
			return nil, nil, fmt.Errorf("method %q is read-only", method)
		}
		return nil, nil, fmt.Errorf("method %q is not read-only", method)
	}
	return service, serviceMethod, nil
}

// isExported returns true of a string is an exported (upper case) name.
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// isExportedOrBuiltin returns true if a type is exported or a builtin.
func isExportedOrBuiltin(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}
