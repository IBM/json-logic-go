package jsonlogic

import (
	"fmt"
)

type operation func(args ...interface{}) (interface{}, error)

var operations = make(map[string]operation)

func opCustom(opName string, arg interface{}, data interface{}) (interface{}, error) {
	if op, ok := operations[opName]; ok {
		switch arg.(type) {
		case []interface{}:
			var err error
			argValues := make([]interface{}, len(arg.([]interface{})))
			for i, argExpr := range arg.([]interface{}) {
				argValues[i], err = applyInterfaces(argExpr, data)
				if err != nil {
					return nil, err
				}
			}
			return op(argValues...)
		default:
			args, err := applyInterfaces(arg, data)
			if err != nil {
				return nil, err
			}
			return op(args)
		}

	}

	return nil, fmt.Errorf("Unknown uperation: %s", opName)
}

// AddOperation allows you to add a custom operation that will run a Go function.
// The `implementation` must be a function with signature `func(args ...interface{}) (interface{}, error)`.
// The first argument is the value(s) passed to the operation, the second argument is an optional data array.
func AddOperation(name string, implementation operation) error {
	if _, ok := operations[name]; ok {
		return fmt.Errorf("Operation exists: %s", name)
	}

	operations[name] = implementation

	return nil
}
