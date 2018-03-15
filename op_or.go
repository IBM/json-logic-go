package jsonlogic

import "fmt"

func opOr(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var lastValue interface{}
	var err error
	for _, e := range valuearray {

		lastValue, err = applyInterfaces(e, data)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
		if truthy(lastValue) {
			return lastValue, nil
		}
	}
	return lastValue, nil
}
