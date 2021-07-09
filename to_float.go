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
	"reflect"
	"strconv"
)

//TODO: See who uses this, and consider the error handling
func interfaceToFloat(value interface{}) float64 {
	metaValue := reflect.ValueOf(value)

	switch metaValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(metaValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(metaValue.Uint())
	case reflect.Float32, reflect.Float64:
		return metaValue.Float()
	case reflect.String:
		floatVal, _ := strconv.ParseFloat(metaValue.String(), 64)
		return floatVal
	}

	return 0
}

func isNumeric(value interface{}) bool {
	switch value.(type) {
	case float32, float64, int, int8, int16, int32, int64, uint, uint16, uint32, uint64, uint8:
		return true
	case string:
		_, err := strconv.ParseFloat(value.(string), 64)
		return err == nil
	default:
		return false
	}
}
