package jsonlogic

func opMissing(value interface{}, data interface{}) interface{} {
	var resultArray = []interface{}{}
	processedValue := applyInterfaces(value, data)

	switch processedValue.(type) {
	case []interface{}:
		valuearray := processedValue.([]interface{})

		for _, e := range valuearray {
			if dataLookup(data, e, nil) == nil {
				resultArray = append(resultArray, e)
			}
		}

	default: //single value
		if dataLookup(data, processedValue, nil) == nil {
			resultArray = append(resultArray, processedValue)
		}
	}
	return resultArray

}
