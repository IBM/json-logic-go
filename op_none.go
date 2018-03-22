package jsonlogic

import "fmt"

func opNone(value interface{}, data interface{}) (interface{}, error) {
	var rule, inputs interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for None operator")
	}
	if len(valuearray) > 0 {
		inputs, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rule = valuearray[1]
		}
	}
	switch inputs.(type) {
	case []interface{}:
		for _, input := range inputs.([]interface{}) {
			lastValue, err := ApplyJSONInterfaces(rule, input)
			if err != nil {
				return nil, err
			}
			ok, err := truthy(lastValue)
			if err != nil {
				return nil, err
			}
			if ok {
				return false, nil
			}
		}
		return true, nil
	default:
		return nil, fmt.Errorf("invalid input for None operator")
	}

}
