package config

import (
	"reflect"
	"strings"
	"testing"

	"github.com/loomnetwork/go-loom"

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
	err := MockSetConfigSetting(config, "MockAppStoreConfig.MockData1", "-50")
	require.NoError(err)
	require.Equal(config.MockAppStoreConfig.MockData1, int64(-50))
	// Set non-exist setting, expect error
	err = MockSetConfigSetting(config, "MockAppStoreConfig.Unon", "50")
	require.Equal(ErrSettingNotFound, err)
	err = MockSetConfigSetting(config, "asbcd", "50")
	require.Equal(ErrSettingNotFound, err)
	// Set negative value to uint64, expect error
	err = MockSetConfigSetting(config, "MockAppStoreConfig.MockData2", "-50")
	require.Equal(ErrInvalidSettingType, err)
	// Set positive value to uint64, expect value change
	err = MockSetConfigSetting(config, "MockAppStoreConfig.MockData2", "50")
	require.Equal(config.MockAppStoreConfig.MockData2, uint64(50))
	// Set string to string, expect value change
	err = MockSetConfigSetting(config, "MockAppStoreConfig.MockData3", "stringgggggg")
	require.Equal(config.MockAppStoreConfig.MockData3, "stringgggggg")
	// Set int to BigUInt, expect value change
	err = MockSetConfigSetting(config, "MockAppStoreConfig.MockData4", "5555555")
	expectedValue := &types.BigUInt{Value: *loom.NewBigUIntFromInt(5555555)}
	require.Equal(config.MockAppStoreConfig.MockData4.String(), expectedValue.String())
	// Set int to pointer, expect error
	err = MockSetConfigSetting(config, "MockAppStoreConfig", "5555555")
	require.Equal(ErrInvalidSettingType, err)
}

type MockAppStoreConfig struct {
	MockData1 int64
	MockData2 uint64
	MockData3 string
	MockData4 *types.BigUInt
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
		},
	}
}

// SetConfigSetting sets value to config field
func MockSetConfigSetting(config *MockConfig, key, value string) error {
	fieldNames := strings.Split(key, ".")
	if len(fieldNames) > 2 {
		return ErrSettingNotFound
	}
	var field reflect.Value
	if len(fieldNames) == 1 {
		configInterface := reflect.ValueOf(config)
		field = reflect.Indirect(configInterface).FieldByName(fieldNames[0])
		if !field.IsValid() {
			return ErrSettingNotFound
		}
	} else if len(fieldNames) == 2 {
		configInterface := reflect.ValueOf(config)
		structInterface := reflect.Indirect(configInterface).FieldByName(fieldNames[0])
		if !structInterface.IsValid() {
			return ErrSettingNotFound
		}
		field = reflect.Indirect(structInterface).FieldByName(fieldNames[1])
		if !field.IsValid() {
			return ErrSettingNotFound
		}
	}
	return setField(&field, value)
}
