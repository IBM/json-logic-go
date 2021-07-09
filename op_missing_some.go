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

func opMissingSome(value interface{}, data interface{}) (interface{}, error) {
	var valuearray []interface{}
	var input interface{}
	var targetCount int
	processedValue, err := ApplyJSONInterfaces(value, data)
	if err != nil {
		return nil, err
	}

	switch processedValue.(type) {
	case []interface{}:
		valuearray = processedValue.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Missing operator")
	}

	if len(valuearray) > 0 {
		switch valuearray[0].(type) {
		case float64:
			targetCount = int(valuearray[0].(float64))
		case float32:
			targetCount = int(valuearray[0].(float32))
		case int:
			targetCount = valuearray[0].(int)
		default:
			return nil, fmt.Errorf("invalid input for Missing operator")
		}

		if len(valuearray) > 1 {
			input = valuearray[1]
		}
	}
	found := 0

	var resultArray = []interface{}{}

	switch input.(type) {
	case []interface{}:
		inputarray := input.([]interface{})

		for _, e := range inputarray {
			res, _ := dataLookup(data, e, nil)
			if res == nil {
				resultArray = append(resultArray, e)
			} else {
				found++
			}
		}

	default: //single value
		res, _ := dataLookup(data, input, nil)
		if res == nil {
			resultArray = append(resultArray, input)
		} else {
			found++
		}
	}

	if found < targetCount {
		return resultArray, nil
	}
	return []interface{}{}, nil

}
