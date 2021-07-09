/*
MIT License

Copyright (c) 2018 IBM

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
		values, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return interfaceToFloat(values), nil
	}
	lastValue, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(lastValue)

	if len(valuearray) == 2 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		return (val + secVal), nil

	} else if len(valuearray) == 3 {
		// sec value
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		// third value
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
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
		values, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return interfaceToFloat(values), nil
	}

	firstValue, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		return (val * secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
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
		values, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return interfaceToFloat(values), nil
	}
	firstValue, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstValue)
	if len(valuearray) == 2 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)
		return (val - secVal), nil
	} else if len(valuearray) == 3 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
		if err != nil {
			return nil, err
		}
		secVal := interfaceToFloat(secValue)

		thirdValue, err := ApplyJSONInterfaces(valuearray[2], data)
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
		values, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return values, nil
	}
	firstVal, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := interfaceToFloat(firstVal)
	if len(valuearray) == 2 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
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
		values, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return values, nil
	}

	firstVal, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	val := int(interfaceToFloat(firstVal))
	if len(valuearray) == 2 {
		secValue, err := ApplyJSONInterfaces(valuearray[1], data)
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
		val, err := ApplyJSONInterfaces(valuearray[0], data)
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
		val, err := ApplyJSONInterfaces(valuearray[0], data)
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
