package jsonlogic

func opMap(value interface{}, data interface{}) []interface{} {
	valuearray := value.([]interface{})
	arrayName := valuearray[0].(map[string]interface{})["var"]
	operation := valuearray[1].(map[string]interface{})
	var result []interface{}
	if data != nil {
		mapData := data.(map[string]interface{})
		array := mapData[arrayName.(string)]
		for _, val := range array.([]interface{}) {
			res := applyInterfaces(operation, val)
			result = append(result, res)
		}
	} else {
		result = []interface{}{}
	}
	return result
}
