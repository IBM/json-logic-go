package jsonlogic

func opMap(value interface{}, data interface{}) []interface{} {
	valuearray := value.([]interface{})
	array := applyInterfaces(valuearray[0], data)
	operation := valuearray[1].(map[string]interface{})
	var result []interface{}

	if array == nil || len(array.([]interface{})) == 0 {
		result = []interface{}{}
	} else {
		for _, val := range array.([]interface{}) {
			res := applyInterfaces(operation, val)
			result = append(result, res)
		}
	}
	return result
}
