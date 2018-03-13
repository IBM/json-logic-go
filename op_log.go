package jsonlogic

import "log"

func opLog(value interface{}) interface{} {
	log.Println(value)
	return value
}
