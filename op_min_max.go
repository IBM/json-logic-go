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

		val := applyInterfaces(valuearray[0], data)

		if !isNumeric(val) {
			return nil
		}

		result := interfaceToFloat(val)

		var floatVal float64

		for i := 1; i < len(valuearray); i++ {
			val = applyInterfaces(valuearray[i], data)

			if !isNumeric(val) {
				return nil
			}

			floatVal = interfaceToFloat(val)

			if floatVal > result {
				result = floatVal
			}
		}

		return result

	default: // A single value
		return applyInterfaces(value, data)
	}
}
