package jsonlogic

func opEqual(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) == interfaceToFloat(rightValue), nil
	}
	return leftValue == rightValue, nil
}

func opEqualStrict(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}
	return leftValue == rightValue, nil
}

func opNotEqual(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) != interfaceToFloat(rightValue), nil
	}
	return leftValue != rightValue, nil
}

func opNotEqualStrict(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}

	return leftValue != rightValue, nil
}

func opSmallerThan(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)

	if len(valuearray) == 3 {
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val < secVal) && (secVal < thirdVal), nil
	}
	return (val < secVal), nil

}

func opGreaterThan(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val > secVal) && (secVal > thirdVal), nil
	}
	return (val > secVal), nil
}

func opSmallerEqThan(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val <= secVal) && (secVal <= thirdVal), nil
	}
	return (val <= secVal), nil
}

func opGreaterEqThan(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	leftValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	rightValue, err := applyInterfaces(valuearray[1], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(leftValue)
	secVal := interfaceToFloat(rightValue)
	if len(valuearray) == 3 {
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val >= secVal) && (secVal >= thirdVal), nil
	}
	return (val >= secVal), nil
}
