package config

import (
	"encoding/json"
	"testing"

	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (t *ConfigTestSuite) SetupTest() {
}

func (t *ConfigTestSuite) TestSetConfigSetting() {
	require := t.Require()
	config := &cctypes.Config{
		AppStoreConfig: &cctypes.AppStoreConfig{},
	}
	err := SetConfigSetting(config, "AppStoreConfig.NumEvmKeysToPrune", "50")
	require.NoError(err)
	require.Equal(config.AppStoreConfig.NumEvmKeysToPrune, uint64(50))
	err = SetConfig(config, "ABC.NumEvmKeysToPrune", "50")
	require.Equal(ErrConfigNotFound, err)
	err = SetConfig(config, "asbcd", "50")
	require.Equal(ErrConfigNotFound, err)
	err = SetConfig(config, "AppStoreConfig.NumEvmKeysToPrune", "true")
	require.Equal(ErrConfigWrongType, err)
}

func (t *ConfigTestSuite) TestStructConvertion() {
	require := t.Require()
	configProtobuf := &cctypes.Config{
		AppStoreConfig: &cctypes.AppStoreConfig{
			NumEvmKeysToPrune: 50,
		},
	}
	str, err := json.Marshal(configProtobuf)
	require.NoError(err)
	var config Config
	err = json.Unmarshal(str, &config)
	require.NoError(err)
	require.Equal(uint64(50), config.AppStoreConfig.NumEvmKeysToPrune)
}
