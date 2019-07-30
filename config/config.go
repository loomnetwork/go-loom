package config

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

type Config struct {
	ConfigProtobuf *cctypes.Config
}

func NewConfig(configProtobuf *cctypes.Config) *Config {
	return &Config{
		ConfigProtobuf: configProtobuf,
	}
}

func SetConfig(config *Config, key, value string) error {
	fieldNames := strings.Split(key, ".")
	if len(fieldNames) > 2 {
		return ErrConfigNotFound
	}
	var field reflect.Value
	if len(fieldNames) == 1 {
		cfgInterface := reflect.ValueOf(config.ConfigProtobuf)
		field = reflect.Indirect(cfgInterface).FieldByName(fieldNames[0])
	} else if len(fieldNames) == 2 {
		cfgInterface := reflect.ValueOf(config.ConfigProtobuf)
		structInterface := reflect.Indirect(cfgInterface).FieldByName(fieldNames[0])
		field = reflect.Indirect(structInterface).FieldByName(fieldNames[1])
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
