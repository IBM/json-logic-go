package jsonlogic

func opReduce(value interface{}, data interface{}) float64 {

	valuearray := value.([]interface{})
	array := applyInterfaces(valuearray[0], data)
	operation := valuearray[1]
	accumulator := valuearray[2].(float64)

	if array == nil {
		return accumulator
	}

	var dat = map[string]interface{}{}
	for _, val := range array.([]interface{}) {
		//  "current" : this element of the array
		// "accumulator" : progress so far, or the initial value
		dat["current"] = val
		dat["accumulator"] = accumulator
		accumulator = applyInterfaces(operation, dat).(float64)
	}

	return accumulator
}
