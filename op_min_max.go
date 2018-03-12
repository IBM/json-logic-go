package jsonlogic

import "math"

func opMax(value interface{}, data interface{}) interface{} {
	if value == nil {
		return 0
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil
		}

		result := applyInterfaces(valuearray[0], data).(float64)

		for expr := range valuearray {
			var val = applyInterfaces(expr, data)

			result = math.Max(result, val.(float64))
		}

		return result

	default: // A single value
		return applyInterfaces(value, data)
	}
}
