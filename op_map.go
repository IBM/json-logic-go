package jsonlogic

func opMap(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	if len(valuearray) == 0 {
		return []interface{}{}, nil
	}
	array, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	var operation interface{}
	if len(valuearray) == 2 {
		operation = valuearray[1].(map[string]interface{})
	}

	var result []interface{}
	if array == nil || len(array.([]interface{})) == 0 {
		result = []interface{}{}
	} else {
		for _, val := range array.([]interface{}) {
			res, err := ApplyJSONInterfaces(operation, val)
			if err != nil && val != nil {
				return nil, err
			}
			result = append(result, res)
		}
	}
	return result, nil
}
