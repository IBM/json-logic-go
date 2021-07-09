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

func opMissing(value interface{}, data interface{}) (interface{}, error) {
	var resultArray = []interface{}{}
	processedValue, err := ApplyJSONInterfaces(value, data)
	if err != nil {
		return nil, err
	}

	switch processedValue.(type) {
	case []interface{}:
		valuearray := processedValue.([]interface{})

		for _, e := range valuearray {
			res, _ := dataLookup(data, e, nil)
			if res == nil {
				resultArray = append(resultArray, e)
			}
		}

	default: //single value
		res, _ := dataLookup(data, processedValue, nil)
		if res == nil {
			resultArray = append(resultArray, processedValue)
		}
	}
	return resultArray, nil

}
