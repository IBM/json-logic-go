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
			res, _ := dataLookup(data, e, nil)
			if res == nil {
				resultArray = append(resultArray, e)
			}
		}

	default: //single value
		res, _ := dataLookup(data, processedValue, nil)
		if res == nil {
			resultArray = append(resultArray, processedValue)
		}
	}
	return resultArray, nil

}
