package jsonlogic

import "fmt"

func opAnd(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for AND operator")
	}

	var lastValue interface{}
	var err error
	for _, e := range valuearray {
		lastValue, err = ApplyJSONInterfaces(e, data)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(lastValue)
		if err != nil {
			return nil, err
		}
		if !ok {
			return lastValue, nil
		}
	}
	return lastValue, nil
}
