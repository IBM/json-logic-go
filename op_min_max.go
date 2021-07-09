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

func opMax(value interface{}, data interface{}) (interface{}, error) {
	_, max, err := findMinMax(value, data)
	if err != nil {
		return nil, err
	}
	return max, nil
}

func opMin(value interface{}, data interface{}) (interface{}, error) {
	min, _, err := findMinMax(value, data)
	if err != nil {
		return nil, err
	}
	return min, nil
}

func findMinMax(value interface{}, data interface{}) (min interface{}, max interface{}, err error) {
	if value == nil {
		return 0, 0, nil
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil, nil, nil
		}

		val, err := ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, nil, err
		}
		if !isNumeric(val) {
			return nil, nil, nil
		}

		min := interfaceToFloat(val)
		max := min

		var floatVal float64

		for i := 1; i < len(valuearray); i++ {
			val, err = ApplyJSONInterfaces(valuearray[i], data)
			if err != nil {
				return nil, nil, err
			}
			if !isNumeric(val) {
				return nil, nil, nil
			}

			floatVal = interfaceToFloat(val)

			if floatVal > max {
				max = floatVal
			} else {
				min = floatVal
			}
		}

		return min, max, nil

	default: // A single value
		val, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, nil, err
		}
		if !isNumeric(val) {
			return nil, nil, nil
		}

		floatVal := interfaceToFloat(val)

		return floatVal, floatVal, nil
	}
}
