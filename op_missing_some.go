package jsonlogic

import "fmt"

func opMissingSome(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	var input interface{}
	var targetCount int
	processedValue, err := ApplyJSONInterfaces(value, data)
	if err != nil {
		return nil, err
	}

	switch processedValue.(type) {
	case []interface{}:
		valuearray = processedValue.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Missing operator")
	}

	if len(valuearray) > 0 {
		switch valuearray[0].(type) {
		case float64:
			targetCount = int(valuearray[0].(float64))
		case float32:
			targetCount = int(valuearray[0].(float32))
		case int:
			targetCount = valuearray[0].(int)
		default:
			return nil, fmt.Errorf("invalid input for Missing operator")
		}

		if len(valuearray) > 1 {
			input = valuearray[1]
		}
	}
	found := 0

	var resultArray = []interface{}{}

	switch input.(type) {
	case []interface{}:
		inputarray := input.([]interface{})

		for _, e := range inputarray {
			res, _ := dataLookup(data, e, nil)
			if res == nil {
				resultArray = append(resultArray, e)
			} else {
				found++
			}
		}

	default: //single value
		res, _ := dataLookup(data, input, nil)
		if res == nil {
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
