package jsonlogic

import "fmt"

func opMap(value interface{}, data interface{}) ([]interface{}, error) {
	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	operation := valuearray[1].(map[string]interface{})
	var result []interface{}

	if array == nil || len(array.([]interface{})) == 0 {
		result = []interface{}{}
	} else {
		for _, val := range array.([]interface{}) {
			res, err := applyInterfaces(operation, val)
			if err != nil {
				return nil, fmt.Errorf("error")
			}
			result = append(result, res)
		}
	}
	return result, nil
}
