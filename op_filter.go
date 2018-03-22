package jsonlogic

import "fmt"

func opFilter(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	var array, operation interface{}
	var err error
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Filter operator")
	}
	if len(valuearray) > 0 {
		array, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			operation = valuearray[1]
		}
	}

	var result = []interface{}{}
	switch array.(type) {
	case []interface{}:
		if array == nil || len(array.([]interface{})) == 0 {
			return result, nil
		}
	default:
		return nil, fmt.Errorf("invalid input for Filter operator")
	}

	for _, val := range array.([]interface{}) {
		res, err := ApplyJSONInterfaces(operation, val)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(res)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, val)
		}
	}

	return result, nil
}
