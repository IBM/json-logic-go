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

func opVar(value interface{}, data interface{}) (interface{}, error) {
	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})
		if len(valuearray) > 0 {
			value1, err := ApplyJSONInterfaces(valuearray[0], data)
			if err != nil {
				return nil, err
			}
			var value2 interface{}
			if len(valuearray) > 1 {
				value2, err = ApplyJSONInterfaces(valuearray[1], data)
				if err != nil {
					return nil, err
				}
			}
			return dataLookup(data, value1, value2)
		}
		//Expected behavior for empty array is to return the entire data
		return data, nil

	default: // A single value
		res, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		result, err := dataLookup(data, res, nil)
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, fmt.Errorf("Element not found: %s", res)
		}
		return result, nil
	}
}
