package jsonlogic

import "fmt"

func opAll(value interface{}, data interface{}) (interface{}, error) {
	var rule, inputs interface{}
	var valuearray []interface{}
	var err error
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for All operator")
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
		if len(inputs.([]interface{})) > 0 {
			for _, input := range inputs.([]interface{}) {
				value, err := ApplyJSONInterfaces(rule, input)
				if err != nil {
					return nil, err
				}
				ok, err := truthy(value)
				if err != nil {
					return nil, err
				}
				if !ok {
					return false, nil
				}

			}
			return true, nil
		}
		return false, nil
	default:
		return nil, fmt.Errorf("invalid input for All operator")
	}

}
