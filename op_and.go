package jsonlogic

func opAnd(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	for _, e := range valuearray {
		if applyInterfaces(e, data) == false {
			return false
		}
	}
	return true
}
