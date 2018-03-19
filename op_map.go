package jsonlogic

func opMap(value interface{}, data interface{}) ([]interface{}, error) {
	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)

	var result []interface{}

	if err != nil {
		result = []interface{}{}
		return result, err
	}
	operation := valuearray[1].(map[string]interface{})

	if array == nil || len(array.([]interface{})) == 0 {
		result = []interface{}{}
	} else {
		for _, val := range array.([]interface{}) {
			res, _ := applyInterfaces(operation, val)
			result = append(result, res)
		}
	}
	return result, nil
}
