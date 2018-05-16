package loom

import (
	"bytes"
	"sort"

	"github.com/loomnetwork/go-loom/types"
)

type (
	Validator    = types.Validator
	ValidatorSet map[string]*Validator
)

func (vs ValidatorSet) Get(pubKey []byte) *Validator {
	return vs[string(pubKey)]
}

func (vs ValidatorSet) Set(v *Validator) {
	vs[string(v.PubKey)] = v
}

func (vs ValidatorSet) Slice() []*Validator {
	vals := make([]*Validator, 0, len(vs))

	for _, v := range vs {
		vals = append(vals, v)
	}

	sort.Sort(validatorsByAddress(vals))
	return vals
}

type validatorsByAddress []*Validator

func (s validatorsByAddress) Len() int {
	return len(s)
}

func (s validatorsByAddress) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s validatorsByAddress) Less(i, j int) bool {
	return bytes.Compare(s[i].PubKey, s[j].PubKey) < 0
}

func NewValidatorSet(vals ...*Validator) ValidatorSet {
	vs := make(ValidatorSet)
	for _, v := range vals {
		vs.Set(v)
	}

	return vs
}
