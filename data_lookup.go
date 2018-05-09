package jsonlogic

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func dataLookup(data interface{}, index interface{}, defaultValue interface{}) (interface{}, error) {
	if index == nil { //Special case.
		return data, nil
	}

	switch index.(type) {
	case string:
		if strings.ContainsRune(index.(string), '.') {
			// Handle dot-notation
			indexArray := strings.Split(index.(string), ".")
			if len(indexArray) > 1 { //TODO: What would "a." mean?
				firstIndex, indexArray := indexArray[0], indexArray[1:]
				retrievedValue, err := dataLookup(data, firstIndex, defaultValue)
				if err != nil {
					return nil, err
				}
				return dataLookup(retrievedValue, strings.Join(indexArray, "."), defaultValue)
			}

		} else if len(index.(string)) == 0 {
			// Special case. See "all", "some", "none"
			return data, nil
		} else if isNumeric(index) {
			// Go to numeric/array
			return dataLookup(data, interfaceToFloat(index), defaultValue)
		}
		// Regular mapping

		var dataVal = reflect.ValueOf(data)

		if !dataVal.IsValid() {
			return defaultValue, nil
		}

		var value reflect.Value

		switch dataVal.Kind() {
		case reflect.Map:
			value = dataVal.MapIndex(reflect.ValueOf(index.(string)))
		case reflect.Struct:
			value = dataVal.FieldByName(index.(string))
		default:
			return nil, errors.Errorf("Unsupported type: %v", dataVal)
		}

		if !value.IsValid() {
			return defaultValue, nil
		}

		switch value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Bool, reflect.Int16, reflect.Int32, reflect.Int64:
			return value.Int(), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return value.Uint(), nil
		case reflect.Float32, reflect.Float64:
			return value.Float(), nil
		case reflect.String:
			return value.String(), nil
		case reflect.Map, reflect.Struct, reflect.Interface, reflect.Array, reflect.Slice:
			return value.Interface(), nil
		default:
			return nil, errors.Errorf("Unsupported type: %v", value)
		}

	case float64:
		// Array index
		metaData := reflect.ValueOf(data)
		switch metaData.Kind() {
		case reflect.Slice, reflect.Array:
			i := int(index.(float64))
			if i >= 0 && i < metaData.Len() {
				return getValue(metaData.Index(i)), nil
			}
		}

	}
	return defaultValue, nil

}
