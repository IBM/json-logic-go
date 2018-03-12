package jsonlogic

func opMax(value interface{}, data interface{}) interface{} {
	if value == nil {
		return 0
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil
		}

		result := interfaceToFloat(applyInterfaces(valuearray[0], data))

		for _, expr := range valuearray {
			var val = interfaceToFloat(applyInterfaces(expr, data))

			if val > result {
				result = val
			}
		}

		return result

	default: // A single value
		return applyInterfaces(value, data)
	}
}
