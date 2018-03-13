package jsonlogic

import (
	"fmt"
)

type operation func(args ...interface{}) interface{}

var operations map[string]operation = make(map[string]operation)

func opCustom(opName string, arg interface{}, data interface{}) (interface{}, error) {
	if op, ok := operations[opName]; ok {
		argValue := applyInterfaces(arg, data)

		switch argValue.(type) {
		case []interface{}:
			return op(argValue.([]interface{})...), nil
		default:
			return op(argValue), nil
		}
	}

	return nil, fmt.Errorf("Unknown uperation: %s", opName)
}

func AddOperation(name string, implementation operation) error {
	if _, ok := operations[name]; ok {
		return fmt.Errorf("Operation exists: %s", name)
	}

	operations[name] = implementation

	return nil
}
