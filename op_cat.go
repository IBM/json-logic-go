package jsonlogic

import (
	"fmt"
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
				res, err := applyInterfaces(val, data)
				if err != nil {
					return nil, fmt.Errorf("error")
				}
				result += res.(string)
			}
		}
	default:
		if value != nil {
			res, err := applyInterfaces(value, data)
			if err != nil {
				return nil, fmt.Errorf("error")
			}
			result = res.(string)
		}
	}
	return result, nil

}
