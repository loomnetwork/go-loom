package config

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
)

var (
	// ErrSettingNotFound indicates that a config does not exist
	ErrSettingNotFound = errors.New("[Application] config not found")
	// ErrInvalidSettingType returned when types of value and config variable mismatch
	ErrInvalidSettingType = errors.New("[Application] wrong variable type")
)

func DefaultConfig() *cctypes.Config {
	return &cctypes.Config{
		AppStore: &cctypes.AppStore{
			NumEvmKeysToPrune: 50,
		},
	}
}

// SetConfigSetting sets value to config field
func SetConfigSetting(config *cctypes.Config, key, value string) error {
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

func setField(field *reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Uint64:
		val, err := strconv.ParseUint(value, 10, 64)
		if err != nil || val < 0 {
			return ErrInvalidSettingType
		}
		field.SetUint(uint64(val))
	case reflect.Int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return ErrInvalidSettingType
		}
		field.SetInt(val)
	case reflect.Bool:
		val, err := strconv.ParseBool(value)
		if err != nil {
			return ErrInvalidSettingType
		}
		field.SetBool(val)
	default:
		return ErrInvalidSettingType
	}
	return nil
}
