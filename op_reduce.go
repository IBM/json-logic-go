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

func opReduce(value interface{}, data interface{}) (interface{}, error) {
	var err error
	var valuearray []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray = value.([]interface{})
	default:
		return nil, fmt.Errorf("invalid input for Reduce operator")
	}

	if len(valuearray) == 0 {
		return nil, nil
	}
	array, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	var operation interface{}
	var accumulator interface{}
	if len(valuearray) == 2 {
		operation = valuearray[1]
	} else if len(valuearray) == 3 {
		operation = valuearray[1]
		accumulator = valuearray[2]
	}

	if err != nil {
		return accumulator, err
	}

	if array == nil {
		return accumulator, nil
	}

	var dat = map[string]interface{}{}
	for _, val := range array.([]interface{}) {
		//  "current" : this element of the array
		// "accumulator" : progress so far, or the initial value
		dat["current"] = val
		dat["accumulator"] = accumulator
		accumulator, err = ApplyJSONInterfaces(operation, dat)
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}
