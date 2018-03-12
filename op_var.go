package jsonlogic

func opVar(value interface{}, data interface{}) interface{} {
	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})
		if len(valuearray) > 0 {
			value1 := applyInterfaces(valuearray[0], data)
			var value2 interface{}
			if len(valuearray) > 1 {
				value2 = applyInterfaces(valuearray[1], data)
			}
			return dataLookup(data, value1, value2)
		}
		//TODO: Expected behavior for empty array?
		return false

	default: // A single value
		return dataLookup(data, applyInterfaces(value, data), nil)
	}
}
