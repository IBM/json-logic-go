package jsonlogic

func opReduce(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, err
	}
	var operation interface{}
	var accumulator interface{}
	if len(valuearray) == 2 {
		operation = valuearray[1]
	} else if len(valuearray) == 3 {
		operation = valuearray[1]
		accumulator = valuearray[2]
	}

	if err != nil {
		return accumulator, err
	}

	if array == nil {
		return accumulator, nil
	}

	var dat = map[string]interface{}{}
	for _, val := range array.([]interface{}) {
		//  "current" : this element of the array
		// "accumulator" : progress so far, or the initial value
		dat["current"] = val
		dat["accumulator"] = accumulator
		accumulator, err = applyInterfaces(operation, dat)
		if err != nil {
			return nil, err
		}
	}

	return accumulator, nil
}
