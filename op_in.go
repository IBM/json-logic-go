package jsonlogic

import (
	"strings"
)

func opIn(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var needle, haystack interface{}
	var err error
	//check needle type
	switch valuearray[0].(type) {
	case string:
		needle = valuearray[0].(string)
	default:
		needle, err = applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
	}

	haystack, err = applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}

	switch haystack.(type) {
	case string:
		return strings.Contains(haystack.(string), needle.(string)), nil
	case []interface{}:
		for _, word := range haystack.([]interface{}) {
			switch word.(type) {
			case string:
				if needle.(string) == word.(string) {
					return true, nil
				}
			// TODO: What if it is a numerical value or something else that is not a string?
			default:
				wordString, err := applyInterfaces(word, data)
				if err != nil {
					return nil, err
				}
				if needle.(string) == wordString.(string) {
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
