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

import (
	"fmt"
	"strconv"
)

func opCat(value interface{}, data interface{}) (interface{}, error) {
	var result string
	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})
		for _, val := range valuearray {
			switch val.(type) {
			case string:
				result += val.(string)
			case int:
				result += strconv.Itoa((val.(int)))
			case float64:
				result += strconv.Itoa((int(val.(float64))))
			default:
				res, err := ApplyJSONInterfaces(val, data)
				if err != nil {
					return nil, err
				}
				switch res.(type) {
				case string:
					result += res.(string)
				default:
					return nil, fmt.Errorf("invalid input for CAT operator")
				}

			}
		}
	default:
		if value != nil {
			res, err := ApplyJSONInterfaces(value, data)
			if err != nil {
				return nil, err
			}
			switch res.(type) {
			case string:
				result = res.(string)
			default:
				return nil, fmt.Errorf("invalid input for CAT operator")
			}
		}
	}
	return result, nil

}
