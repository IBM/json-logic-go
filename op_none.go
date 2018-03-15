package jsonlogic

import "fmt"

func opNone(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	inputs, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	rule := valuearray[1]

	if len(inputs.([]interface{})) > 0 {
		for _, input := range inputs.([]interface{}) {
			lastValue, err := applyInterfaces(rule, input)
			if err != nil {
				return nil, fmt.Errorf("error")
			}
			if truthy(lastValue) {
				return false, nil
			}
		}
		return true, nil
	}
	return true, nil
}
