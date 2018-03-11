package jsonlogic

func opOr(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	var lastValue interface{}
	for _, e := range valuearray {
		lastValue = applyInterfaces(e, data)
		if truthy(lastValue) {
			return lastValue
		}
	}
	return lastValue
}
