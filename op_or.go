package jsonlogic

func opOr(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var lastValue interface{}
	var err error
	for _, e := range valuearray {

		lastValue, err = applyInterfaces(e, data)
		if err != nil {
			return nil, err
		}
		if truthy(lastValue) {
			return lastValue, nil
		}
	}
	return lastValue, nil
}
