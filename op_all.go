package jsonlogic

func opAll(value interface{}, data interface{}) interface{} {
	valuearray := value.([]interface{})
	inputs := applyInterfaces(valuearray[0], data)
	rule := valuearray[1]

	if len(inputs.([]interface{})) > 0 {
		result := true
		for _, input := range inputs.([]interface{}) {
			result = result && truthy(applyInterfaces(rule, input))
		}
		return result
	}
	return false
}
