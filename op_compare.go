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
