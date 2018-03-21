package jsonlogic

import (
	"fmt"
)

func opSum(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		values, err := applyInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return interfaceToFloat(values), nil
	}
	lastValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(lastValue)

	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		return (val + secVal), nil

	} else if len(valuearray) == 3 {
		// sec value
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		// third value
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		return interfaceToFloat(values), nil
	}

	firstValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		return (val * secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		return interfaceToFloat(values), nil
	}
	firstValue, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		return (val - secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		thirdValue, err := applyInterfaces(valuearray[2], data)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		return values, nil
	}
	firstVal, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstVal)
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		if secVal == float64(0) {
			return nil, fmt.Errorf("divide by Zero")
		}
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
			return nil, err
		}
		return values, nil
	}

	firstVal, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := int(interfaceToFloat(firstVal))
	if len(valuearray) == 2 {
		secValue, err := applyInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		ok, err := truthy(val)
		if err != nil {
			return nil, err
		}
		return !ok, nil
	default:
		ok, err := truthy(value)
		if err != nil {
			return nil, err
		}
		return !ok, nil
	}
}

func opDoubleNot(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
		val, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(val)
		if err != nil {
			return nil, err
		}
		return ok, nil
	default:
		ok, err := truthy(value)
		if err != nil {
			return nil, err
		}
		return ok, nil
	}
}
