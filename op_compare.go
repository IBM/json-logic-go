package jsonlogic

func opEqual(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) == interfaceToFloat(rightValue), nil
	}
	return leftValue == rightValue, nil
}

func opEqualStrict(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}
	return leftValue == rightValue, nil
}

func opNotEqual(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) != interfaceToFloat(rightValue), nil
	}
	return leftValue != rightValue, nil
}

func opNotEqualStrict(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}

	return leftValue != rightValue, nil
}

func opSmallerThan(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)

	if len(valuearray) == 3 {
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val < secVal) && (secVal < thirdVal), nil
	}
	return (val < secVal), nil

}

func opGreaterThan(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}

	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val > secVal) && (secVal > thirdVal), nil
	}
	return (val > secVal), nil
}

func opSmallerEqThan(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}
	}

	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val <= secVal) && (secVal <= thirdVal), nil
	}
	return (val <= secVal), nil
}

func opGreaterEqThan(value interface{}, data interface{}) (interface{}, error) {
	var leftValue, rightValue interface{}
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	case interface{}:
		leftValue = value
	}

	if len(valuearray) > 0 {
		leftValue, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		if len(valuearray) > 1 {
			rightValue, err = ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
		}
	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val >= secVal) && (secVal >= thirdVal), nil
	}
	return (val >= secVal), nil
}
