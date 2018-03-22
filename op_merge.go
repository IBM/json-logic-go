package jsonlogic

func opMerge(value interface{}, data interface{}) (interface{}, error) {
	var result []interface{}
	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})
		if len(valuearray) == 0 {
			result = value.([]interface{})
		}
		for _, val := range valuearray {
			processedValue, err := ApplyJSONInterfaces(val, data)
			if err != nil {
				return nil, err
			}
			switch processedValue.(type) {
			case []interface{}:
				res, err := opMerge(processedValue, data)
				if err != nil {
					return nil, err
				}
				result = append(result, res.([]interface{})...)
			case interface{}:
				result = append(result, processedValue)
			}
		}
	default:
		if value != nil {
			res, err := ApplyJSONInterfaces(value, data)
			if err != nil {
				return nil, err
			}
			result = append(result, res)
		}
	}

	return result, nil
}
