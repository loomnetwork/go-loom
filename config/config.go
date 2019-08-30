package config

import (
	"errors"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/loomnetwork/go-loom"

	cctypes "github.com/loomnetwork/go-loom/builtin/types/chainconfig"
	"github.com/loomnetwork/go-loom/types"
)

var (
	// ErrSettingNotFound indicates that a setting does not exist
	ErrSettingNotFound = errors.New("[Application] setting not found")
	// ErrInvalidSettingType returned when types of value and setting variable mismatch
	ErrInvalidSettingType = errors.New("[Application] wrong variable type")
)

// DefaultConfig returns an on-chain config instance with all reference fields initialized to
// non-nil values (this is needed for the reflection code in SetConfigSetting to work, for now...)
func DefaultConfig() *cctypes.Config {
	return &cctypes.Config{
		AppStore:     &cctypes.AppStoreConfig{},
		Evm:          &cctypes.EvmConfig{},
		NonceHandler: &cctypes.NonceHandlerConfig{},
	}
}

// SetConfigSetting sets value of a config field.
// The key can consist of one or two elements, e.g. "fieldA", "sectionA.fieldB".
// The value string should be parsable into a string, integer, or a BigUInt.
func SetConfigSetting(config interface{}, key, value string) error {
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
		if err != nil {
			return ErrInvalidSettingType
		}
		field.SetUint(val)
	case reflect.Int64:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return ErrInvalidSettingType
		}
		field.SetInt(val)
	case reflect.Ptr:
		bigIntAmount := big.NewInt(0)
		if _, ok := bigIntAmount.SetString(value, 0); !ok {
			return ErrInvalidSettingType
		}
		bigUintAmount := loom.NewBigUInt(bigIntAmount)
		protoBigUint := &types.BigUInt{Value: *bigUintAmount}
		if field.Elem().Type() != reflect.TypeOf(*protoBigUint) {
			return ErrInvalidSettingType
		}
		field.Elem().Set(reflect.ValueOf(*protoBigUint))
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
