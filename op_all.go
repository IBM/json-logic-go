package jsonlogic

func opAll(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	inputs, _ := ApplyJSONInterfaces(valuearray[0], data)
	rule := valuearray[1]

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
}
