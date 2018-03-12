package jsonlogic

import (
	"strconv"
)

//TODO: See who uses this, and consider the error handling
func interfaceToFloat(value interface{}) float64 {
	if isNumeric(value) {
		var val float64
		switch value.(type) {
		case float64:
			val = value.(float64)
		case int:
			val = float64(value.(int))
		case string:
			val, _ = strconv.ParseFloat(value.(string), 64)
		}
		return val
	}

	return 0
}

func isNumeric(value interface{}) bool {
	switch value.(type) {
	case float32, float64, int, int8, int16, int32, int64:
		return true
	case string:
		_, err := strconv.ParseFloat(value.(string), 64)
		return err == nil
	default:
		return false
	}
}
