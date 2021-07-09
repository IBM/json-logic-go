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

func opOr(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Or operator")
	}
	var lastValue interface{}
	for _, e := range valuearray {

		lastValue, err = ApplyJSONInterfaces(e, data)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(lastValue)
		if err != nil {
			return nil, err
		}
		if ok {
			return lastValue, nil
		}
	}
	return lastValue, nil
}
