package jsonlogic

func opMap(value interface{}, data interface{}) []interface{} {
	valuearray := value.([]interface{})
	array := applyInterfaces(valuearray[0], data)
	operation := valuearray[1].(map[string]interface{})
	var result []interface{}

	if data != nil {
		mapData := data.(map[string]interface{})
		switch array.(type) {
		case string: //get from Data
			arrays := mapData[array.(string)]
			for _, val := range arrays.([]interface{}) {
				res := applyInterfaces(operation, val)
				result = append(result, res)
			}
		default: // Hardcoded
			for _, val := range array.([]interface{}) {
				res := applyInterfaces(operation, val)
				result = append(result, res)
			}
		}
	} else {
		if array == nil || len(array.([]interface{})) == 0 {
			result = []interface{}{}
		} else { // Hardcoded
			for _, val := range array.([]interface{}) {
				res := applyInterfaces(operation, val)
				result = append(result, res)
			}
		}
	}
	return result
}
