package config

import (
	"testing"

	"github.com/loomnetwork/go-loom"

	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	"github.com/loomnetwork/go-loom/types"
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
	config := MockDefaultConfig()
	// Set int64, expect negative value set
	err := SetConfigSetting(config, "MockAppStoreConfig.MockData1", "-50")
	require.NoError(err)
	require.Equal(config.MockAppStoreConfig.MockData1, int64(-50))
	// Set non-exist setting, expect error
	err = SetConfigSetting(config, "MockAppStoreConfig.Unon", "50")
	require.Equal(ErrSettingNotFound, err)
	err = SetConfigSetting(config, "asbcd", "50")
	require.Equal(ErrSettingNotFound, err)
	// Set negative value to uint64, expect error
	err = SetConfigSetting(config, "MockAppStoreConfig.MockData2", "-50")
	require.Equal(ErrInvalidSettingType, err)
	// Set positive value to uint64, expect value change
	SetConfigSetting(config, "MockAppStoreConfig.MockData2", "50")
	require.Equal(config.MockAppStoreConfig.MockData2, uint64(50))
	// Set string to string, expect value change
	SetConfigSetting(config, "MockAppStoreConfig.MockData3", "stringgggggg")
	require.Equal(config.MockAppStoreConfig.MockData3, "stringgggggg")
	// Set int to BigUInt, expect value change
	SetConfigSetting(config, "MockAppStoreConfig.MockData4", "5555555")
	expectedValue := &types.BigUInt{Value: *loom.NewBigUIntFromInt(5555555)}
	require.Equal(config.MockAppStoreConfig.MockData4.String(), expectedValue.String())
	// Set int to pointer, expect error
	err = SetConfigSetting(config, "MockAppStoreConfig", "5555555")
	require.Equal(ErrInvalidSettingType, err)

	// Set random string to bool, expect error
	err = SetConfigSetting(config, "MockAppStoreConfig.MockData5", "abcde")
	require.Equal(ErrInvalidSettingType, err)
	// Set number to bool, expect error
	err = SetConfigSetting(config, "MockAppStoreConfig.MockData5", "12")
	require.Equal(ErrInvalidSettingType, err)
	// Set true to bool, expect true
	SetConfigSetting(config, "MockAppStoreConfig.MockData5", "true")
	require.Equal(config.MockAppStoreConfig.MockData5, true)
	// Set false to bool, expect false
	SetConfigSetting(config, "MockAppStoreConfig.MockData5", "false")
	require.Equal(config.MockAppStoreConfig.MockData5, false)
}

func (t *ConfigTestSuite) TestNilConfigSetting() {
	require := t.Require()
	var cfg cctypes.Config
	require.Nil(cfg.GetAppStore())
	require.Nil(cfg.GetEvm())
	cfg2 := DefaultConfig()
	require.NotNil(cfg2.GetAppStore())
	require.NotNil(cfg2.GetEvm())
}

type MockAppStoreConfig struct {
	MockData1 int64
	MockData2 uint64
	MockData3 string
	MockData4 *types.BigUInt
	MockData5 bool
}

type MockConfig struct {
	MockAppStoreConfig *MockAppStoreConfig
}

func MockDefaultConfig() *MockConfig {
	return &MockConfig{
		MockAppStoreConfig: &MockAppStoreConfig{
			MockData1: int64(-1),
			MockData2: uint64(2),
			MockData3: "string_data",
			MockData4: &types.BigUInt{Value: *loom.NewBigUIntFromInt(3)},
			MockData5: true,
		},
	}
}
