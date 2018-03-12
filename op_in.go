package jsonlogic

import (
	"strings"
)

func opIn(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	var subString interface{}
	//check subString type
	switch valuearray[0].(type) {
	case string:
		subString = valuearray[0].(string)
	default:
		subString = applyInterfaces(valuearray[0], data)
	}

	switch valuearray[1].(type) {
	case string:
		valuearray = value.([]interface{})
		return strings.Contains(valuearray[1].(string), subString.(string))
	case []interface{}:
		for _, word := range valuearray[1].([]interface{}) {
			switch word.(type) {
			case string:
				if subString.(string) == word.(string) {
					return true
				}
			default:
				wordString := applyInterfaces(word, data)
				if subString.(string) == wordString.(string) {
					return true
				}
				return false
			}

		}
	default:
		return false
	}
	return false
}
