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

func opSome(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	var rule, inputs interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Some operator")
	}
	if len(valuearray) == 0 {
		return nil, nil
	}
	inputs, err = ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	if len(valuearray) > 1 {
		rule = valuearray[1]
	}

	if len(inputs.([]interface{})) > 0 {
		for _, input := range inputs.([]interface{}) {
			res, err := ApplyJSONInterfaces(rule, input)
			if err != nil {
				return nil, err
			}
			ok, err := truthy(res)
			if err != nil {
				return nil, err
			}
			if ok {
				return true, nil
			}
		}
		return false, nil
	}
	return false, nil
}
