// NOTE this file was taken from https://github.com/gorilla/rpc/blob/master/map.go
// modified highly

// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package contractpb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/gogo/protobuf/proto"
)

var (
	// Precompute the reflect.Type of some types
	typeOfError         = reflect.TypeOf((*error)(nil)).Elem()
	typeOfContext       = reflect.TypeOf((*Context)(nil)).Elem()
	typeOfStaticContext = reflect.TypeOf((*StaticContext)(nil)).Elem()
	typeOfPBMessage     = reflect.TypeOf((*proto.Message)(nil)).Elem()
)

var (
	ErrServiceNotFound = errors.New("service not found")
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

type methodSig int

const (
	methodSigUnknown methodSig = iota
	methodSigInit
	methodSigCall
	methodSigStaticCall
)

type serviceMethod struct {
	method    reflect.Method // receiver method
	argsType  reflect.Type   // type of the request argument
	methodSig methodSig
}

// ----------------------------------------------------------------------------
// serviceMap
// ----------------------------------------------------------------------------

// serviceMap is a registry for services.
type serviceMap struct {
	mutex    sync.Mutex
	services map[string]*service
}

func detMethodSig(method reflect.Method) (methodSig, error) {
	mtype := method.Type

	// Method must be exported.
	if method.PkgPath != "" {
		return methodSigUnknown, errors.New("method is not exported")
	}

	// Method needs four ins: receiver, plugin.Context, *args.
	if mtype.NumIn() != 3 {
		return methodSigUnknown, errors.New("method does not have correct number of args")
	}

	n := mtype.NumOut()
	switch {
	case n == 1:
		firstRet := mtype.Out(0)
		if !firstRet.Implements(typeOfPBMessage) && !firstRet.Implements(typeOfError) {
			return methodSigUnknown, errors.New("return value must be proto.Message or error")
		}
	case n == 2:
		firstRet := mtype.Out(0)
		secondRet := mtype.Out(1)
		if !firstRet.Implements(typeOfPBMessage) || !secondRet.Implements(typeOfError) {
			return methodSigUnknown, errors.New("return value must be proto.Message, error")
		}
	case n > 2:
		return methodSigInit, errors.New("methods must have at most 2 return values")
	}

	contextType := mtype.In(1)
	args := mtype.In(2)

	// Second argument must be a pointer and must be something that implements the
	// plugin.[Static]Context interface
	if !contextType.Implements(typeOfContext) && !contextType.Implements(typeOfStaticContext) {
		return methodSigUnknown, errors.New("methods must take in a context as the first argument")
	}

	if !args.Implements(typeOfPBMessage) {
		return methodSigUnknown, errors.New("the second argument must be a proto.Message")
	}

	if method.Name == "Init" {
		if !contextType.Implements(typeOfContext) {
			return methodSigInit, errors.New("init does not take in a context")
		}

		if mtype.NumOut() != 1 {
			return methodSigInit, errors.New("init must have a single return")
		}

		return methodSigInit, nil
	}

	if contextType.Implements(typeOfContext) {
		return methodSigCall, nil
	}

	return methodSigStaticCall, nil
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
		if method.Name == "Meta" {
			continue
		}
		methodSig, err := detMethodSig(method)
		if err != nil {
			continue
		}
		srvMethod := &serviceMethod{
			method:    method,
			argsType:  method.Type.In(2).Elem(),
			methodSig: methodSig,
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
func (m *serviceMap) Get(method string) (*service, *serviceMethod, error) {
	parts := strings.Split(method, ".")
	if len(parts) != 2 {
		err := fmt.Errorf("service/method request ill-formed: %q", method)
		return nil, nil, err
	}
	m.mutex.Lock()
	service := m.services[parts[0]]
	m.mutex.Unlock()
	if service == nil {
		return nil, nil, ErrServiceNotFound
	}
	serviceMethod := service.methods[parts[1]]
	if serviceMethod == nil {
		return nil, nil, fmt.Errorf("contract method '%s' not found", method)
	}
	return service, serviceMethod, nil
}

// isExported returns true of a string is an exported (upper case) name.
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}
