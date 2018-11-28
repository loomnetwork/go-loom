package plugin

import (
	"bytes"
	"testing"

	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/util"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	addr1 := loom.MustParseAddress("chain:0xb16a379ec18d4093666f8f38b11a3071c920207d")

	s := CreateFakeContext(addr1, addr1)

	s.Set([]byte("bob1"), []byte("asasdfasdf"))
	s.Set([]byte("123321123"), []byte("asasdfasdf"))
	s.Set([]byte("bob4"), []byte("asasdfasdf"))
	s.Set([]byte("bob5"), []byte("asasdfasdf"))
	s.Set([]byte("bob6"), []byte("asasdfasdf"))
	s.Set([]byte("afsddsf"), []byte("asasdfasdf"))

	data := s.Range([]byte("bob"))

	assert.Equal(t, 4, len(data))

	//The mock context uses map underneath and the real context does not so ordering will be different then real server!
	//	assert.Equal(t, string(s.makeKey([]byte("bob5"))), string(data[0].Key))
}

func TestPrefixedKeys(t *testing.T) {
	addr1 := loom.MustParseAddress("chain:0xb16a379ec18d4093666f8f38b11a3071c920207d")

	c := CreateFakeContext(addr1, addr1)
	prefix := []byte("my prefix")
	unprefixedKey := []byte("placeholder")
	noContextKey := util.PrefixKey(prefix, unprefixedKey)

	// key is c.address + prefix + unprefixedKey
	assert.Equal(t, 0, bytes.Compare(c.recoverKey(c.makeKey(noContextKey), prefix), unprefixedKey))
}
