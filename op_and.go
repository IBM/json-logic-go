package jsonlogic

func opAnd(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	var lastValue interface{}
	for _, e := range valuearray {
		if !truthy(applyInterfaces(e, data)) {
			return e
		}
		lastValue = e
	}
	return lastValue
}
