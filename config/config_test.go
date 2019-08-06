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
		AppStore: &cctypes.AppStore{},
	}
	err := SetConfigSetting(config, "AppStore.NumEvmKeysToPrune", "50")
	require.NoError(err)
	require.Equal(config.AppStore.NumEvmKeysToPrune, uint64(50))
	err = SetConfigSetting(config, "ABC.NumEvmKeysToPrune", "50")
	require.Equal(ErrSettingNotFound, err)
	err = SetConfigSetting(config, "asbcd", "50")
	require.Equal(ErrSettingNotFound, err)
	err = SetConfigSetting(config, "AppStore.NumEvmKeysToPrune", "true")
	require.Equal(ErrInvalidSettingType, err)
}

func (t *ConfigTestSuite) TestStructConvertion() {
	require := t.Require()
	configProtobuf := &cctypes.Config{
		AppStore: &cctypes.AppStore{
			NumEvmKeysToPrune: 50,
		},
	}
	str, err := json.Marshal(configProtobuf)
	require.NoError(err)
	var config Config
	err = json.Unmarshal(str, &config)
	require.NoError(err)
	require.Equal(uint64(50), config.AppStore.NumEvmKeysToPrune)
}
