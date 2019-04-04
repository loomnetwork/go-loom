package contractpb

import (
	"testing"

	"github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/stretchr/testify/suite"
)

type ContractContextTestSuite struct {
	suite.Suite
	addr1 loom.Address
	addr2 loom.Address
}

func TestContractContextSuite(t *testing.T) {
	suite.Run(t, new(ContractContextTestSuite))
}

func (s *ContractContextTestSuite) SetupTest() {
	s.addr1 = loom.MustParseAddress("chain:0xb16a379ec18d4093666f8f38b11a3071c920207d")
	s.addr2 = loom.MustParseAddress("chain:0xfa4c7920accfd66b86f5fd0e69682a79f762d49e")
}

func (s *ContractContextTestSuite) TestPermissions() {
	require := s.Require()
	callerAddr := s.addr1
	contractAddr := s.addr2
	context := plugin.CreateFakeContext(callerAddr, contractAddr)
	ctx := WrapPluginContext(context)

	perm1 := []byte("perm1")
	perm2 := []byte("perm2")
	role1 := "role1"
	role2 := "role2"
	roles := []string{role1, role2}

	hasPerm, _,_ := ctx.HasPermissionFor(callerAddr, perm1, roles)
	require.False(hasPerm)

	ctx.GrantPermissionTo(callerAddr, perm1, role1)
	hasPerm, _,_= ctx.HasPermissionFor(callerAddr, perm1, []string{role1})
	require.True(hasPerm)
	hasPerm,_,_ = ctx.HasPermissionFor(callerAddr, perm1, []string{role2})
	require.False(hasPerm)

	ctx.GrantPermissionTo(callerAddr, perm2, role1)
	hasPerm, _,_ = ctx.HasPermissionFor(callerAddr, perm2, []string{role1})
	require.True(hasPerm)
	hasPerm,_,_ = ctx.HasPermissionFor(callerAddr, perm2, []string{role2})
	require.False(hasPerm)

	ctx.RevokePermissionFrom(callerAddr, perm1, role1)
	hasPerm, _,_ = ctx.HasPermissionFor(callerAddr, perm1, []string{role1})
	require.False(hasPerm)
	hasPerm,_,_ = ctx.HasPermissionFor(callerAddr, perm2, []string{role1})
	require.True(hasPerm)
}
