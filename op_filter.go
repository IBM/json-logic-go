package jsonlogic

func opFilter(value interface{}, data interface{}) ([]interface{}, error) {
	valuearray := value.([]interface{})
	array, err := ApplyJSONInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	operation := valuearray[1]
	var result = []interface{}{}

	if array == nil || len(array.([]interface{})) == 0 {
		return result, nil
	}

	for _, val := range array.([]interface{}) {
		res, err := ApplyJSONInterfaces(operation, val)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(res)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, val)
		}
	}

	return result, nil
}
