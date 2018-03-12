package jsonlogic

func opSum(value interface{}, data interface{}) float64 {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return interfaceToFloat(applyInterfaces(value, data))
	}

	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	if len(valuearray) == 2 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		return (val + secVal)
	} else if len(valuearray) == 3 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val + secVal + thirdVal)
	}

	return val

}

func opMult(value interface{}, data interface{}) float64 {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return interfaceToFloat(applyInterfaces(value, data))
	}

	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	if len(valuearray) == 2 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		return (val * secVal)
	} else if len(valuearray) == 3 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val * secVal * thirdVal)
	}

	return val
}

func opSub(value interface{}, data interface{}) float64 {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return interfaceToFloat(applyInterfaces(value, data))
	}

	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	if len(valuearray) == 2 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		return (val - secVal)
	} else if len(valuearray) == 3 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		thirdVal := interfaceToFloat(applyInterfaces(valuearray[2], data))
		return (val - secVal - thirdVal)
	}

	return (-val)
}

func opDiv(value interface{}, data interface{}) float64 {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return interfaceToFloat(applyInterfaces(value, data))
	}

	val := interfaceToFloat(applyInterfaces(valuearray[0], data))
	if len(valuearray) == 2 {
		secVal := interfaceToFloat(applyInterfaces(valuearray[1], data))
		return (val / secVal)
	}

	return val
}

func opMod(value interface{}, data interface{}) float64 {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return interfaceToFloat(applyInterfaces(value, data))
	}

	val := int(interfaceToFloat(applyInterfaces(valuearray[0], data)))
	if len(valuearray) == 2 {
		secVal := int(interfaceToFloat(applyInterfaces(valuearray[1], data)))
		return float64(val % secVal)
	}

	return float64(val)
}

func opNot(value interface{}, data interface{}) bool {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
		return !(truthy(applyInterfaces(valuearray[0], data)))
	default:
		return !(truthy(value))
	}
}

func opDoubleNot(value interface{}, data interface{}) bool {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
		return truthy(applyInterfaces(valuearray[0], data))
	default:
		return truthy(value)
	}
}
