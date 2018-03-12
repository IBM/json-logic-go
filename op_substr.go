package jsonlogic

func opSubstr(value interface{}, data interface{}) string {
	valuearray := value.([]interface{})
	var subString interface{}
	switch valuearray[0].(type) {
	case string:
		subString = valuearray[0].(string)
	default:
		subString = applyInterfaces(valuearray[0], data)
	}

	firstIndex := int(valuearray[1].(float64))
	if firstIndex < 0 {
		firstIndex = len(subString.(string)) + firstIndex
	}
	if len(valuearray) == 3 {
		secIndex := int(valuearray[2].(float64))
		if secIndex < 0 {
			secIndex = len(subString.(string)) + secIndex
			return subString.(string)[firstIndex:secIndex]
		}
		return subString.(string)[firstIndex : firstIndex+secIndex]
	}
	return subString.(string)[firstIndex:]

}
