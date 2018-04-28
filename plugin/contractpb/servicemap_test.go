package contractpb

import (
	"fmt"
	"testing"
)

type mockPB struct {
}

func (m *mockPB) Reset() {
}

func (m *mockPB) String() string {
	return ""
}

func (m *mockPB) ProtoMessage() {
}

type fakeTxPB struct {
	mockPB
}

type fakeResponsePB struct {
	mockPB
}

type fakeResponse struct {
}

type fakeContract struct {
}

type FakeServiceMapContext interface{}

// These methods SHOULD NOT be auto-registered:
func (c *fakeContract) IgnoredMethod1() {}
func (c *fakeContract) ignoredMethod2() {}
func (c *fakeContract) IgnoredMethod3(ctx Context, tx *fakeTxPB) int {
	return 0
}

// This method is ignored because the return type is  not a pb
func (c *fakeContract) IgnoredMethod4(ctx Context, tx *fakeTxPB) (*fakeResponse, error) {
	return nil, nil
}

// This method is ignored because the first argument is not a plugin context
func (c *fakeContract) IgnoredMethod5(ctx FakeServiceMapContext, tx *fakeTxPB) error {
	return nil
}

// Ditto
func (c *fakeContract) IgnoredMethod6(ctx FakeServiceMapContext, tx *fakeTxPB) (*fakeResponse, error) {
	return nil, nil
}

// These methods SHOULD be auto-registered
func (c *fakeContract) TxHandler1(ctx Context, tx *fakeTxPB) error {
	return nil
}

func (c *fakeContract) QueryHandler1(ctx Context, tx *fakeTxPB) (*fakeResponsePB, error) {
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
		if _, _, err := srvMap.Get(methodName); err == nil {
			t.Errorf("Error: %s should not be registered", methodName)
		}
		if _, _, err := srvMap.Get(methodName); err == nil {
			t.Errorf("Error: %s should not be registered", methodName)
		}
	}

	if _, _, err := srvMap.Get("fakeContract.TxHandler1"); err != nil {
		t.Errorf("Error: fakeContract.TxHandler1 should be registered")
	}

	if _, _, err := srvMap.Get("fakeContract.QueryHandler1"); err != nil {
		t.Errorf("Error: fakeContract.QueryHandler1 should be registered")
	}
}
