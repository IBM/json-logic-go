package jsonlogic

func opMissing(value interface{}, data interface{}) (interface{}, error) {
	var resultArray = []interface{}{}
	processedValue, err := applyInterfaces(value, data)
	if err != nil {
		return nil, err
	}

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
	return resultArray, nil

}
