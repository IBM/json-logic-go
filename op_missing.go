package jsonlogic

func opMissing(value interface{}, data interface{}) interface{} {
	var resultArray = []interface{}{}

	switch value.(type) {
	case []interface{}:
		valuearray := value.([]interface{})

		for _, e := range valuearray {
			if dataLookup(data, e, nil) == nil {
				resultArray = append(resultArray, e)
			}
		}

	default: //single value
		if dataLookup(data, value, nil) == nil {
			resultArray = append(resultArray, value)
		}
	}
	return resultArray

}
