package jsonlogic

import (
	"strconv"
)

func opCat(value interface{}, data interface{}) (interface{}, error) {
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
				res, err := ApplyJSONInterfaces(val, data)
				if err != nil {
					return nil, err
				}
				result += res.(string)
			}
		}
	default:
		if value != nil {
			res, err := ApplyJSONInterfaces(value, data)
			if err != nil {
				return nil, err
			}
			result = res.(string)
		}
	}
	return result, nil

}
