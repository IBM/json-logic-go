package jsonlogic

func dataLookup(data interface{}, index interface{}, defaultValue interface{}) interface{} {
	switch data.(type) {
	case map[string]interface{}:
		switch index.(type) {
		case string:
			value, ok := data.(map[string]interface{})[index.(string)]
			if ok == true {
				return value
			}
			return defaultValue

		}
	case []interface{}:
		dataArray := data.([]interface{})
		switch index.(type) {
		case float64:
			indexInt := int(index.(float64))
			if len(dataArray) >= indexInt+1 {
				return dataArray[indexInt]
			}
			return defaultValue
		}
	case float64: //See "all", "some", "none"
		if index == "" {
			return data
		}
		return defaultValue
	}
	return defaultValue
}
