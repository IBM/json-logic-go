package jsonlogic

import "fmt"

func opSome(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	var rule, inputs interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Some operator")
	}
	if len(valuearray) == 0 {
		return nil, nil
	}
	inputs, err = ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	if len(valuearray) > 1 {
		rule = valuearray[1]
	}

	if len(inputs.([]interface{})) > 0 {
		for _, input := range inputs.([]interface{}) {
			res, err := ApplyJSONInterfaces(rule, input)
			if err != nil {
				return nil, err
			}
			ok, err := truthy(res)
			if err != nil {
				return nil, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
	return false, nil
}
