package jsonlogic

func opFilter(value interface{}, data interface{}) []interface{} {
	valuearray := value.([]interface{})
	array := applyInterfaces(valuearray[0], data)
	operation := valuearray[1]
	var result = []interface{}{}

	if array == nil || len(array.([]interface{})) == 0 {
		return result
	}

	for _, val := range array.([]interface{}) {
		res := applyInterfaces(operation, val)
		if truthy(res) {
			result = append(result, val)
		}
	}

	return result
}
