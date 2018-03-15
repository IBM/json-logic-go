package jsonlogic

import "fmt"

func opReduce(value interface{}, data interface{}) (interface{}, error) {

	valuearray := value.([]interface{})
	array, err := applyInterfaces(valuearray[0], data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	operation := valuearray[1]
	accumulator := valuearray[2]

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
		fmt.Println("accum", accumulator)
		if err != nil {
			return nil, fmt.Errorf("error")
		}
	}

	return accumulator, nil
}
