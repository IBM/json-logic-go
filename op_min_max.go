package jsonlogic

func opMax(value interface{}, data interface{}) (interface{}, error) {
	_, max := findMinMax(value, data)

	return max, nil
}

func opMin(value interface{}, data interface{}) (interface{}, error) {
	min, _ := findMinMax(value, data)

	return min, nil
}

func findMinMax(value interface{}, data interface{}) (min interface{}, max interface{}) {
	if value == nil {
		return 0, 0
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil, nil
		}

		val, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if !isNumeric(val) {
			return nil, nil
		}

		min := interfaceToFloat(val)
		max := min

		var floatVal float64

		for i := 1; i < len(valuearray); i++ {
			val, err = applyInterfaces(valuearray[i], data)
			if err != nil {
				return nil, err
			}
			if !isNumeric(val) {
				return nil, nil
			}

			floatVal = interfaceToFloat(val)

			if floatVal > max {
				max = floatVal
			} else {
				min = floatVal
			}
		}

		return min, max

	default: // A single value
		val, err := applyInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		if !isNumeric(val) {
			return nil, nil
		}

		floatVal := interfaceToFloat(val)

		return floatVal, floatVal
	}
}
