package jsonlogic

import "fmt"

func opSum(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return interfaceToFloat(values), nil
	}
	lastValue, err := applyInterfaces(valuearray[0], data)
	val := interfaceToFloat(lastValue)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)

		return (val + secVal), nil

	} else if len(valuearray) == 3 {
		// sec value
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)

		// third value
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		thirdVal := interfaceToFloat(thirdValue)

		return (val + secVal + thirdVal), nil
	}

	return val, nil

}

func opMult(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return interfaceToFloat(values), nil
	}

	firstValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)
		return (val * secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val * secVal * thirdVal), nil
	}

	return val, nil
}

func opSub(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return interfaceToFloat(values), nil
	}
	firstValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)
		return (val - secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)

		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		thirdVal := interfaceToFloat(thirdValue)
		return (val - secVal - thirdVal), nil
	}

	return (-val), nil
}

func opDiv(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return values, nil
	}
	firstVal, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	val := interfaceToFloat(firstVal)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := interfaceToFloat(secValue)
		return (val / secVal), nil
	}

	return val, nil
}

func opMod(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return values, nil
	}

	firstVal, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	val := int(interfaceToFloat(firstVal))
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		secVal := int(interfaceToFloat(secValue))
		return float64(val % secVal), nil
	}

	return float64(val), nil
}

func opNot(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
		val, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return !(truthy(val)), nil
	default:
		return !(truthy(value)), nil
	}
}

func opDoubleNot(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
		val, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		return truthy(val), nil
	default:
		return truthy(value), nil
	}
}
