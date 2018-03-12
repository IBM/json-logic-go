package jsonlogic

import "fmt"

func opLog(value interface{}) interface{} {
	fmt.Println(value)
	return value
}
