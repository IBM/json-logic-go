package jsonlogic

import "log"

func opLog(value interface{}) (interface{}, error) {
	log.Println(value)
	return value, nil
}
