package plugin

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
)

var (
	// ErrConfigNotFound indicates that a config does not exist
	ErrConfigNotFound = errors.New("[Application] config not found")
	// ErrConfigWrongType returned when types of value and config variable mismatch
	ErrConfigWrongType = errors.New("[Application] wrong variable type")
)

var (
	cfgSettingVersion = 1
)

func mockDefaultConfig() *cctypes.Config {
	return &cctypes.Config{
		Version: uint64(cfgSettingVersion),
		AppStoreConfig: &cctypes.AppStoreConfig{
			DeletedVmKeys: 50,
		},
	}
}

func mockSetConfig(config *cctypes.Config, key, value string) error {
	fieldNames := strings.Split(key, ".")
	if len(fieldNames) > 2 {
		return ErrConfigNotFound
	}
	var field reflect.Value
	if len(fieldNames) == 1 {
		cfgInterface := reflect.ValueOf(config)
		field = reflect.Indirect(cfgInterface).FieldByName(fieldNames[0])
	} else if len(fieldNames) == 2 {
		cfgInterface := reflect.ValueOf(config)
		structInterface := reflect.Indirect(cfgInterface).FieldByName(fieldNames[0])
		field = reflect.Indirect(structInterface).FieldByName(fieldNames[1])
	}
	return mockSetField(&field, value)
}

func mockSetField(field *reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Uint64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil || val < 0 {
			return ErrConfigWrongType
		}
		field.SetUint(uint64(val))
	case reflect.Int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return ErrConfigWrongType
		}
		field.SetInt(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(value)
		if err != nil {
			return ErrConfigWrongType
		}
		field.SetBool(val)
	default:
		return ErrConfigWrongType
	}
	return nil
}
