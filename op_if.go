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

func opIf(value interface{}, data interface{}) (interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil, nil
		}

		condition, err := applyInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}

		if len(valuearray) == 1 {
			return condition, nil
		}

		ok, err := truthy(condition)
		if err != nil {
			return nil, err
		}
		if ok {
			res, err := applyInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
			return res, nil
		}

		if len(valuearray) > 2 {
			return opIf(valuearray[2:], data)
		}

		return nil, nil

	default: // A single value
		res, err := applyInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
