package jsonlogic

func opMerge(value interface{}, data interface{}) []interface{} {
	var result []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})
		if len(valuearray) == 0 {
			result = value.([]interface{})
		}
		for _, val := range valuearray {
			processedValue := applyInterfaces(val, data)
			switch processedValue.(type) {
			case []interface{}:
				result = append(result, opMerge(processedValue, data)...)
			case interface{}:
				result = append(result, processedValue)
			}
		}
	default:
		if value != nil {
			result = append(result, applyInterfaces(value, data))
		}
	}

	return result
}
