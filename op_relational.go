package jsonlogic

import (
	"fmt"
	"reflect"
	"strconv"
)

func interFaceToFloat(value interface{}) float64 {
	var val float64
	switch value.(type) {
	case float64:
		val = value.(float64)
	case int:
		val = float64(value.(int))
	case string:
		intVal, err := strconv.Atoi(value.(string))
		if err != nil {
			fmt.Println("unexpected numeric value", value.(string))
			panic(0)
		}
		val = float64(intVal)
	case nil:
		val = 0
	default:
		fmt.Println("unexpected type", reflect.TypeOf(value))
		val = 0
	}

	return val
}

func smallerThan(value interface{}, other interface{}) bool {
	val := interFaceToFloat(value)
	otherVal := interFaceToFloat(other)
	return val < otherVal
}

func greaterThan(value interface{}, other interface{}) bool {
	val := interFaceToFloat(value)
	otherVal := interFaceToFloat(other)
	return val > otherVal
}

func smallerEqualThan(value interface{}, other interface{}) bool {
	val := interFaceToFloat(value)
	otherVal := interFaceToFloat(other)
	return val <= otherVal
}

func greaterEqualThan(value interface{}, other interface{}) bool {
	val := interFaceToFloat(value)
	otherVal := interFaceToFloat(other)
	return val >= otherVal
}
