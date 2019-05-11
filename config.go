package loom

import (
	"fmt"
	"strconv"
)

type Config interface {
	DPOS() *DPOSConfig
	GetConfig(string) string
}

const (
	dposFeeFloor = "dpos.feeFloor"
)

type ChainConfig struct {
	cfg  map[string]string
	dpos *DPOSConfig
}

func NewChainConfig(cfg map[string]string) Config {
	chainConfig := &ChainConfig{}
	chainConfig.dpos = NewDPOSConfig(cfg)
	return chainConfig
}

func (cfg *ChainConfig) DPOS() *DPOSConfig {
	return cfg.dpos
}

func (cfg *ChainConfig) GetConfig(key string) string {
	return cfg.cfg[key]
}

type DPOSConfig struct {
	cfg map[string]string
}

func NewDPOSConfig(cfg map[string]string) *DPOSConfig {
	dposConfig := &DPOSConfig{
		cfg: cfg,
	}
	return dposConfig
}

func (dpos *DPOSConfig) FeeFloor(val int64) int64 {
	feeFloor, err := getInt64(dposFeeFloor, dpos.cfg)
	if err != nil {
		return val
	}
	return feeFloor
}

func getInt64(key string, cfg map[string]string) (int64, error) {
	if value, ok := cfg[key]; ok {
		v, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return v, nil
		}
		return 0, err
	}
	return 0, fmt.Errorf("Key %s not found", key)
}
