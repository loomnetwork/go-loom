package plugin

import (
	"fmt"
	"testing"
)

type FakeTx struct{}

type fakeTx struct{}

type fakeResponse struct{}

type FakeResponse struct{}

type fakeContract struct{}

type FakeContext interface{}

// These methods SHOULD NOT be auto-registered:
func (c *fakeContract) IgnoredMethod1()                        {}
func (c *fakeContract) ignoredMethod2()                        {}
func (c *fakeContract) IgnoredMethod3(ctx Context)             {}
func (c *fakeContract) IgnoredMethod4(ctx Context, tx *FakeTx) {}
func (c *fakeContract) IgnoredMethod5(ctx Context, tx *FakeTx) int {
	return 0
}

// This method will be ignored because the type of the second argument is not exported
func (c *fakeContract) IgnoredMethod6(ctx Context, tx *fakeTx) error {
	return nil
}

// This method is ignored because the return type is not exported
func (c *fakeContract) IgnoredMethod7(ctx Context, tx *FakeTx) (*fakeResponse, error) {
	return nil, nil
}

// This method is ignored because the first argument in not a plugin context
func (c *fakeContract) IgnoredMethod8(ctx FakeContext, tx *FakeTx) error {
	return nil
}

// Ditto
func (c *fakeContract) IgnoredMethod9(ctx FakeContext, tx *FakeTx) (*FakeResponse, error) {
	return nil, nil
}

// These methods SHOULD be auto-registered
func (c *fakeContract) TxHandler1(ctx Context, tx *FakeTx) error {
	return nil
}

func (c *fakeContract) QueryHandler1(ctx Context, tx *FakeTx) (*FakeResponse, error) {
	return nil, nil
}

func TestServiceMapDuplicateServices(t *testing.T) {
	srvMap := new(serviceMap)
	if err := srvMap.Register(&fakeContract{}, "fakeContract"); err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	if err := srvMap.Register(&fakeContract{}, "fakeContract"); err == nil {
		t.Errorf("Error: duplicate service names should not be allowed")
	}
}

func TestServiceMapAutoDiscovery(t *testing.T) {
	srvMap := new(serviceMap)
	if err := srvMap.Register(&fakeContract{}, "fakeContract"); err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	for i := 1; i < 10; i++ {
		methodName := fmt.Sprintf("fakeContract.IgnoredMethod%d", i)
		if _, _, err := srvMap.Get(methodName, false); err == nil {
			t.Errorf("Error: %s should not be registered", methodName)
		}
		if _, _, err := srvMap.Get(methodName, true); err == nil {
			t.Errorf("Error: %s should not be registered", methodName)
		}
	}

	if _, _, err := srvMap.Get("fakeContract.TxHandler1", false); err != nil {
		t.Errorf("Error: fakeContract.TxHandler1 should be registered")
	}

	if _, _, err := srvMap.Get("fakeContract.TxHandler1", true); err == nil {
		t.Errorf("Error: fakeContract.TxHandler1 should not be read-only")
	}

	if _, _, err := srvMap.Get("fakeContract.QueryHandler1", true); err != nil {
		t.Errorf("Error: fakeContract.QueryHandler1 should be registered")
	}

	if _, _, err := srvMap.Get("fakeContract.QueryHandler1", false); err == nil {
		t.Errorf("Error: fakeContract.QueryHandler1 should be read-only")
	}
}
