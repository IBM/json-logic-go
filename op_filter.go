package jsonlogic

func opFilter(value interface{}, data interface{}) ([]interface{}, error) {
	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	operation := valuearray[1]
	var result = []interface{}{}

	if array == nil || len(array.([]interface{})) == 0 {
		return result, nil
	}

	for _, val := range array.([]interface{}) {
		res, err := applyInterfaces(operation, val)
		if err != nil {
			return nil, err
		}
		if truthy(res) {
			result = append(result, val)
		}
	}

	return result, nil
}
