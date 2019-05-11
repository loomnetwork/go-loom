package plugin

import (
	"fmt"
	"strconv"

	loom "github.com/loomnetwork/go-loom"
)

const (
	dposFeeFloor = "dpos.feeFloor"
)

type FakeConfig struct {
	cfg  map[string]string
	dpos loom.DPOSConfig
}

func (fcfg *FakeConfig) DPOS() loom.DPOSConfig {
	return fcfg.dpos
}

func (fcfg *FakeConfig) GetConfig(key string) string {
	return fcfg.cfg[key]
}

type FakeDPOSConfig struct {
	cfg map[string]string
}

func (fdpos *FakeDPOSConfig) FeeFloor(val int64) int64 {
	feeFloor, err := getInt64(dposFeeFloor, fdpos.cfg)
	if err != nil {
		return val
	}
	return feeFloor
}

func getInt64(key string, cfg map[string]string) (int64, error) {
	if value, ok := cfg[dposFeeFloor]; ok {
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return v, nil
		}
		return 0, err
	}
	return 0, fmt.Errorf("Key %s not found", key)
}
