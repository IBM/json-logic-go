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
	"reflect"
)

//Per http://jsonlogic.com/truthy.html

func truthy(input interface{}) (bool, error) {
	switch input.(type) {
	case bool:
		return input.(bool), nil
	case float64:
		return (input.(float64) != 0), nil
	case int:
		return (input.(int) != 0), nil
	case []interface{}: //Real world are interfaces, but tests use real types
		return (len(input.([]interface{})) != 0), nil
	case string:
		return (len(input.(string)) != 0), nil
	case nil:
		return false, nil
	case map[string]interface{}:
		return true, nil
	default:
		return false, fmt.Errorf("truthy unexpected type %v", reflect.TypeOf(input))
	}
}
