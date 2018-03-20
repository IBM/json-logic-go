package jsonlogic

import (
	"fmt"
	"strings"
)

func dataLookup(data interface{}, index interface{}, defaultValue interface{}) (interface{}, error) {
	if index == nil { //Special case.
		return data, nil
	}

	switch index.(type) {
	case string:
		if strings.ContainsRune(index.(string), '.') {
			// Handle dot-notation
			indexArray := strings.Split(index.(string), ".")
			if len(indexArray) > 1 { //TODO: What would "a." mean?
				firstIndex, indexArray := indexArray[0], indexArray[1:]
				retrievedValue, err := dataLookup(data, firstIndex, defaultValue)
				if err != nil {
					return nil, err
				}
				return dataLookup(retrievedValue, strings.Join(indexArray, "."), defaultValue)
			}

		} else if len(index.(string)) == 0 {
			// Special case. See "all", "some", "none"
			return data, nil
		} else if isNumeric(index) {
			// Go to numeric/array
			return dataLookup(data, interfaceToFloat(index), defaultValue)
		}
		// Regular mapping
		switch data.(type) {
		case map[string]interface{}:
			value, ok := data.(map[string]interface{})[index.(string)]
			if ok == true {
				return value, nil
			}

			if defaultValue == nil {
				return nil, fmt.Errorf("error finding value")
			}
		default:
			if defaultValue == nil {
				return nil, fmt.Errorf("error finding value")
			}
		}

	case float64:
		// Array index
		switch data.(type) {
		case []interface{}:
			dataArray := data.([]interface{})
			indexInt := int(index.(float64))
			if len(dataArray) >= indexInt+1 {
				return dataArray[indexInt], nil
			}
		}

	}
	return defaultValue, nil

}
