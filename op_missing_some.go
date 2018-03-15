package jsonlogic

import (
	"fmt"
)

func opMissingSome(value interface{}, data interface{}) (interface{}, error) {
	processedValue, err := applyInterfaces(value, data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	valuearray := processedValue.([]interface{})
	targetCount := int(valuearray[0].(float64))
	found := 0
	input := valuearray[1]

	var resultArray = []interface{}{}

	switch input.(type) {
	case []interface{}:
		inputarray := input.([]interface{})

		for _, e := range inputarray {
			if dataLookup(data, e, nil) == nil {
				resultArray = append(resultArray, e)
			} else {
				found++
			}
		}

	default: //single value
		if dataLookup(data, input, nil) == nil {
			resultArray = append(resultArray, input)
		} else {
			found++
		}
	}

	if found < targetCount {
		return resultArray, nil
	}
	return []interface{}{}, nil

}
