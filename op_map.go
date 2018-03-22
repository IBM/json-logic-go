package jsonlogic

import "fmt"

func opMap(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	var err error
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Map operator")
	}

	if len(valuearray) == 0 {
		return []interface{}{}, nil
	}
	array, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	var operation interface{}
	if len(valuearray) > 1 {
		switch valuearray[1].(type) {
		case map[string]interface{}:
			operation = valuearray[1].(map[string]interface{})
		default:
			return nil, fmt.Errorf("invalid input for Map operator")
		}
	}

	var result []interface{}
	switch array.(type) {
	case []interface{}:
		if array == nil || len(array.([]interface{})) == 0 {
			result = []interface{}{}
			return result, nil
		}
	default:
		return nil, fmt.Errorf("invalid input for Map operator")
	}

	for _, val := range array.([]interface{}) {
		res, err := ApplyJSONInterfaces(operation, val)
		if err != nil && val != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return result, nil
}
