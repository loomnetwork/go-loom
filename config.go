package loom

import (
	"fmt"
	"strconv"
)

type Config interface {
	DPOS() *DPOSConfig
	AppStore() *AppStoreConfig
	GetConfig(string) string
}

const (
	dposFeeFloor          = "dpos.feeFloor"
	appstoreDeletedVMKeys = "appstore.deletedVMKeys"
)

type ChainConfig struct {
	cfg      map[string]string
	dpos     *DPOSConfig
	appstore *AppStoreConfig
}

func NewChainConfig(cfg map[string]string) Config {
	chainConfig := &ChainConfig{
		cfg: cfg,
	}
	chainConfig.dpos = NewDPOSConfig(chainConfig)
	chainConfig.appstore = NewAppStoreConfig(chainConfig)
	return chainConfig
}

func (cfg *ChainConfig) DPOS() *DPOSConfig {
	return cfg.dpos
}

func (cfg *ChainConfig) AppStore() *AppStoreConfig {
	return cfg.appstore
}

func (cfg *ChainConfig) GetConfig(key string) string {
	return cfg.cfg[key]
}

type DPOSConfig struct {
	*ChainConfig
}

func NewDPOSConfig(chainConfig *ChainConfig) *DPOSConfig {
	return &DPOSConfig{
		ChainConfig: chainConfig,
	}
}

func (dpos *DPOSConfig) FeeFloor(val int64) int64 {
	feeFloor, err := getInt64(dposFeeFloor, dpos.cfg)
	if err != nil {
		return val
	}
	return feeFloor
}

type AppStoreConfig struct {
	*ChainConfig
}

func NewAppStoreConfig(chainConfig *ChainConfig) *AppStoreConfig {
	return &AppStoreConfig{
		ChainConfig: chainConfig,
	}
}

func (appstore *AppStoreConfig) DeletedVmKeys(val uint64) uint64 {
	deletedVMKeys, err := getInt64(appstoreDeletedVMKeys, appstore.cfg)
	if err != nil {
		return val
	}
	return uint64(deletedVMKeys)
}

// utility functions
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
