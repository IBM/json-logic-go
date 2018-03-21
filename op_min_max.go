package jsonlogic

func opMax(value interface{}, data interface{}) (interface{}, error) {
	_, max, err := findMinMax(value, data)
	if err != nil {
		return nil, err
	}
	return max, nil
}

func opMin(value interface{}, data interface{}) (interface{}, error) {
	min, _, err := findMinMax(value, data)
	if err != nil {
		return nil, err
	}
	return min, nil
}

func findMinMax(value interface{}, data interface{}) (min interface{}, max interface{}, err error) {
	if value == nil {
		return 0, 0, nil
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil, nil, nil
		}

		val, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, nil, err
		}
		if !isNumeric(val) {
			return nil, nil, nil
		}

		min := interfaceToFloat(val)
		max := min

		var floatVal float64

		for i := 1; i < len(valuearray); i++ {
			val, err = applyInterfaces(valuearray[i], data)
			if err != nil {
				return nil, nil, err
			}
			if !isNumeric(val) {
				return nil, nil, nil
			}

			floatVal = interfaceToFloat(val)

			if floatVal > max {
				max = floatVal
			} else {
				min = floatVal
			}
		}

		return min, max, nil

	default: // A single value
		val, err := applyInterfaces(value, data)
		if err != nil {
			return nil, nil, err
		}
		if !isNumeric(val) {
			return nil, nil, nil
		}

		floatVal := interfaceToFloat(val)

		return floatVal, floatVal, nil
	}
}
