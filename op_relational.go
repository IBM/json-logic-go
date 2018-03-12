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

func opSmallerThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interFaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interFaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interFaceToFloat(applyInterfaces(valuearray[2], data))
		return (val < secVal) && (secVal < thirdVal)
	}
	return (val < secVal)

}

func opGreaterThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interFaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interFaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interFaceToFloat(applyInterfaces(valuearray[2], data))
		return (val > secVal) && (secVal > thirdVal)
	}
	return (val > secVal)
}

func opSmallerEqThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interFaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interFaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interFaceToFloat(applyInterfaces(valuearray[2], data))
		return (val <= secVal) && (secVal <= thirdVal)
	}
	return (val <= secVal)
}

func opGreaterEqThan(value interface{}, data interface{}) bool {
	valuearray := value.([]interface{})
	val := interFaceToFloat(applyInterfaces(valuearray[0], data))
	secVal := interFaceToFloat(applyInterfaces(valuearray[1], data))
	if len(valuearray) == 3 {
		thirdVal := interFaceToFloat(applyInterfaces(valuearray[2], data))
		return (val >= secVal) && (secVal >= thirdVal)
	}
	return (val >= secVal)
}
