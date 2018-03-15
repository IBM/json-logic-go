package jsonlogic

import (
	"fmt"
	"strings"
)

func opIn(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var subString interface{}
	var err error
	//check subString type
	switch valuearray[0].(type) {
	case string:
		subString = valuearray[0].(string)
	default:
		subString, err = applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
	}

	switch valuearray[1].(type) {
	case string:
		valuearray = value.([]interface{})
		return strings.Contains(valuearray[1].(string), subString.(string)), nil
	case []interface{}:
		for _, word := range valuearray[1].([]interface{}) {
			switch word.(type) {
			case string:
				if subString.(string) == word.(string) {
					return true, nil
				}
			default:
				wordString, err := applyInterfaces(word, data)
				if err != nil {
					return nil, fmt.Errorf("error")
				}
				if subString.(string) == wordString.(string) {
					return true, nil
				}
				return false, nil
			}

		}
	default:
		return false, nil
	}
	return false, nil
}
