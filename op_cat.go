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
				res, err := ApplyJSONInterfaces(val, data)
				if err != nil {
					return nil, err
				}
				switch res.(type) {
				case string:
					result += res.(string)
				default:
					return nil, fmt.Errorf("invalid input for CAT operator")
				}

			}
		}
	default:
		if value != nil {
			res, err := ApplyJSONInterfaces(value, data)
			if err != nil {
				return nil, err
			}
			switch res.(type) {
			case string:
				result = res.(string)
			default:
				return nil, fmt.Errorf("invalid input for CAT operator")
			}
		}
	}
	return result, nil

}
