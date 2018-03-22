package jsonlogic

func opAnd(value interface{}, data interface{}) (interface{}, error) {
	valuearray := value.([]interface{})
	var lastValue interface{}
	var err error
	for _, e := range valuearray {
		lastValue, err = ApplyJSONInterfaces(e, data)
		if err != nil {
			return nil, err
		}
		ok, err := truthy(lastValue)
		if err != nil {
			return nil, err
		}
		if !ok {
			return lastValue, nil
		}
	}
	return lastValue, nil
}
