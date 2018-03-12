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
			switch val.(type) {
			case []interface{}:
				result = append(result, opMerge(val, data)...)
			case interface{}:
				result = append(result, val)
			}
		}
	default:
		if value != nil {
			result = append(result, value)
		}
	}

	return result
}
