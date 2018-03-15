package jsonlogic

import (
	"fmt"
)

func opReduce(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	operation := valuearray[1]
	accumulator := valuearray[2]

	//TODO : support dot notation in current variable
	err = parseOperation(operation)
	if err != nil {
		return nil, err
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
		accumulator, err = applyInterfaces(operation, dat)
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}

func parseOperation(operation interface{}) error {
	switch operation.(type) {
	case map[string]interface{}:
		for key, variable := range operation.(map[string]interface{}) {
			switch variable.(type) {
			case string:
				if key == "var" && variable != "current" && variable != "accumulator" {
					return fmt.Errorf("Error: wrong variable for reduce operator : %s", variable)
				}
			default:
				if parseOperation(variable) != nil {
					return parseOperation(variable)
				}
			}
		}
	case []interface{}:
		for _, variable := range operation.([]interface{}) {
			if parseOperation(variable) != nil {
				return parseOperation(variable)
			}

		}
	default:
		return nil
	}
	return nil
}
