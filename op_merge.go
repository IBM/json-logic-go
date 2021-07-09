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

func opMerge(value interface{}, data interface{}) (interface{}, error) {
	var result []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})
		if len(valuearray) == 0 {
			result = value.([]interface{})
		}
		for _, val := range valuearray {
			processedValue, err := ApplyJSONInterfaces(val, data)
			if err != nil {
				return nil, err
			}
			switch processedValue.(type) {
			case []interface{}:
				res, err := opMerge(processedValue, data)
				if err != nil {
					return nil, err
				}
				result = append(result, res.([]interface{})...)
			case interface{}:
				result = append(result, processedValue)
			}
		}
	default:
		if value != nil {
			res, err := ApplyJSONInterfaces(value, data)
			if err != nil {
				return nil, err
			}
			result = append(result, res)
		}
	}

	return result, nil
}
