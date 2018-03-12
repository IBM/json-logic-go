package jsonlogic

import (
	"strings"
)

func dataLookup(data interface{}, index interface{}, defaultValue interface{}) interface{} {
	switch data.(type) {
	case map[string]interface{}:
		switch index.(type) {
		case string:
			indexString := index.(string)
			if strings.ContainsRune(indexString, '.') {
				currentData := data
				var value interface{} = defaultValue
				var ok bool
				for _, e := range strings.Split(indexString, ".") {
					value, ok = currentData.(map[string]interface{})[e]
					if ok == true {
						switch value.(type) {
						case map[string]interface{}:
							currentData = value
						}
					}
				}
				return value
			}

			value, ok := data.(map[string]interface{})[indexString]
			if ok == true {
				return value
			}

			return defaultValue

		}
	case []interface{}:
		dataArray := data.([]interface{})
		switch index.(type) {
		case float64:
			indexInt := int(index.(float64))
			if len(dataArray) >= indexInt+1 {
				return dataArray[indexInt]
			}
			return defaultValue
		case string:
			if isNumeric(index) {
				indexInt := int(interfaceToFloat(index))
				if len(dataArray) >= indexInt+1 {
					return dataArray[indexInt]
				}
				return defaultValue
			}
		}
	case float64: //See "all", "some", "none"
		if index == "" {
			return data
		}
		return defaultValue
	}
	return defaultValue
}
