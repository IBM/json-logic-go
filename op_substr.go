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

import "fmt"

func opSubstr(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	var firstIndex, secIndex int
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Substr operator")
	}

	if len(valuearray) == 0 {
		return nil, nil
	}
	var subString interface{}
	switch valuearray[0].(type) {
	case string:
		subString = valuearray[0].(string)
	default:
		subString, err = ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}
	}
	switch subString.(type) {
	case string:
		break
	default:
		return nil, fmt.Errorf("invalid input for Substr operator")
	}

	if len(valuearray) > 1 {
		switch valuearray[1].(type) {
		case float64:
			firstIndex = int(valuearray[1].(float64))
		case float32:
			firstIndex = int(valuearray[1].(float32))
		case int:
			firstIndex = valuearray[1].(int)
		default:
			return nil, fmt.Errorf("invalid input for Substr operator")
		}
	}
	if firstIndex < 0 {
		switch subString.(type) {
		case string:
			firstIndex = len(subString.(string)) + firstIndex
		}
	}
	if len(valuearray) > 2 {
		switch valuearray[2].(type) {
		case float64:
			secIndex = int(valuearray[2].(float64))
		case float32:
			secIndex = int(valuearray[2].(float32))
		case int:
			secIndex = valuearray[2].(int)
		default:
			return nil, fmt.Errorf("invalid input for Substr operator")
		}

		if secIndex < 0 {
			secIndex = len(subString.(string)) + secIndex
			return subString.(string)[firstIndex:secIndex], nil
		}
		return subString.(string)[firstIndex : firstIndex+secIndex], nil
	}
	return subString.(string)[firstIndex:], nil

}
