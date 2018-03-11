package jsonlogic

// The if statement typically takes 3 arguments: a condition (if), what to do if it’s true (then), and what to do if it’s false (else), like:
// {"if" : [ true, "yes", "no" ]}
//
// If can also take more than 3 arguments, and will pair up arguments like if/then elseif/then elseif/then else. Like:
// {"if" : [
// 	{"<": [{"var":"temp"}, 0] }, "freezing",
// 	{"<": [{"var":"temp"}, 100] }, "liquid",
// 	"gas"
// ]}

func opIf(value interface{}, data interface{}) interface{} {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil
		}

		condition := applyInterfaces(valuearray[0], data)

		if len(valuearray) == 1 {
			return condition
		}

		if truthy(condition) {
			return applyInterfaces(valuearray[1], data)
		}

		if len(valuearray) > 2 {
			return opIf(valuearray[2:], data)
		}

		return nil

	default: // A single value
		return applyInterfaces(value, data)
	}
}

// func toBoolean(val interface{}) bool {
// 	if val == nil {
// 		return false
// 	}

// 	switch val.(type) {
// 	case bool:
// 		return val.(bool)
// 	case string:
// 		return len(val.(string)) > 0
// 	default:
// 		return false
// 	}
// }
