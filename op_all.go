package jsonlogic

func opAll(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	inputs := applyInterfaces(valuearray[0], data)
	rule := valuearray[1]

	if len(inputs.([]interface{})) > 0 {
		for _, input := range inputs.([]interface{}) {
			if truthy(applyInterfaces(rule, input)) == false {
				return false
			}
		}
		return true
	}
	return false
}