package jsonlogic

import "fmt"

func opReduce(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Reduce operator")
	}

	if len(valuearray) == 0 {
		return nil, nil
	}
	array, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	var operation interface{}
	var accumulator interface{}
	if len(valuearray) == 2 {
		operation = valuearray[1]
	} else if len(valuearray) == 3 {
		operation = valuearray[1]
		accumulator = valuearray[2]
	}

	if err != nil {
		return accumulator, err
	}

	if array == nil {
		return accumulator, nil
	}

	var dat = map[string]interface{}{}
	for _, val := range array.([]interface{}) {
		//  "current" : this element of the array
		// "accumulator" : progress so far, or the initial value
		dat["current"] = val
		dat["accumulator"] = accumulator
		accumulator, err = ApplyJSONInterfaces(operation, dat)
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}
