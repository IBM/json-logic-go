package jsonlogic

func opOr(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	for _, e := range valuearray {
		if applyInterfaces(e, data) == true {
			return e
		}
	}
	return false
}
