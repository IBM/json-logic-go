package jsonlogic

func opEqual(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	leftValue := applyInterfaces(valuearray[0], data)
	rightValue := applyInterfaces(valuearray[1], data)

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) == interfaceToFloat(rightValue)
	}
	return leftValue == rightValue
}

func opEqualStrict(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	return applyInterfaces(valuearray[0], data) == applyInterfaces(valuearray[1], data)
}

func opNotEqual(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	leftValue := applyInterfaces(valuearray[0], data)
	rightValue := applyInterfaces(valuearray[1], data)

	if isNumeric(leftValue) && isNumeric(rightValue) {
		return interfaceToFloat(leftValue) != interfaceToFloat(rightValue)
	}
	return leftValue != rightValue
}

func opNotEqualStrict(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	return applyInterfaces(valuearray[0], data) != applyInterfaces(valuearray[1], data)
}

func opSmallerThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val < secVal) && (secVal < thirdVal)
	}
	return (val < secVal)

}

func opGreaterThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val > secVal) && (secVal > thirdVal)
	}
	return (val > secVal)
}

func opSmallerEqThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val <= secVal) && (secVal <= thirdVal)
	}
	return (val <= secVal)
}

func opGreaterEqThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val >= secVal) && (secVal >= thirdVal)
	}
	return (val >= secVal)
}
