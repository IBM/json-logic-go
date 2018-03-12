package jsonlogic

import (
	"strconv"
)

func opCat(value interface{}, data interface{}) string {
	var result string
	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})
		for _, val := range valuearray {
			switch val.(type) {
			case string:
				result += val.(string)
			case int:
				result += strconv.Itoa((val.(int)))
			case float64:
				result += strconv.Itoa((int(val.(float64))))
			default:
				result += applyInterfaces(val, data).(string)
			}
		}
	default:
		if value != nil {
			result = value.(string)
		}
	}
	return result

}
