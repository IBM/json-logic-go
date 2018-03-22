package jsonlogic

import "fmt"

func opVar(value interface{}, data interface{}) (interface{}, error) {
	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})
		if len(valuearray) > 0 {
			value1, err := ApplyJSONInterfaces(valuearray[0], data)
			if err != nil {
				return nil, err
			}
			var value2 interface{}
			if len(valuearray) > 1 {
				value2, err = ApplyJSONInterfaces(valuearray[1], data)
				if err != nil {
					return nil, err
				}
			}
			return dataLookup(data, value1, value2)
		}
		//Expected behavior for empty array is to return the entire data
		return data, nil

	default: // A single value
		res, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		result, err := dataLookup(data, res, nil)
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, fmt.Errorf("element Not found")
		}
		return dataLookup(data, res, nil)
	}
}
