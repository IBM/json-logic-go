package jsonlogic

func opVar(value interface{}, data interface{}) (interface{}, error) {
	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})
		if len(valuearray) > 0 {
			value1, err := applyInterfaces(valuearray[0], data)
			if err != nil {
				return nil, err
			}
			var value2 interface{}
			if len(valuearray) > 1 {
				value2, err = applyInterfaces(valuearray[1], data)
				if err != nil {
					return nil, err
				}
			}
			return dataLookup(data, value1, value2), nil
		}
		//Expected behavior for empty array is to return the entire data
		return data, nil

	default: // A single value
		res, err := applyInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return dataLookup(data, res, nil), nil
	}
}
