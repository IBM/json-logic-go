package jsonlogic

import (
	"reflect"
	"strconv"
)

//TODO: See who uses this, and consider the error handling
func interfaceToFloat(value interface{}) float64 {
	metaValue := reflect.ValueOf(value)

	switch metaValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(metaValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(metaValue.Uint())
	case reflect.Float32, reflect.Float64:
		return metaValue.Float()
	case reflect.String:
		floatVal, _ := strconv.ParseFloat(metaValue.String(), 64)
		return floatVal
	}

	return 0
}

func isNumeric(value interface{}) bool {
	switch value.(type) {
	case float32, float64, int, int8, int16, int32, int64, uint, uint16, uint32, uint64, uint8:
		return true
	case string:
		_, err := strconv.ParseFloat(value.(string), 64)
		return err == nil
	default:
		return false
	}
}
