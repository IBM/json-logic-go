package jsonlogic

import (
	"fmt"
	"reflect"
	"strconv"
)

func interfaceToFloat(value interface{}) float64 {
	var val float64
	switch value.(type) {
	case float64:
		val = value.(float64)
	case int:
		val = float64(value.(int))
	case string:
		tempVal, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			fmt.Println("unexpected numeric value", value.(string))
			panic(0)
		}
		val = tempVal
	case nil:
		val = 0
	default:
		fmt.Println("unexpected type", reflect.TypeOf(value))
		val = 0
	}

	return val
}
