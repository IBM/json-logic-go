package jsonlogic

import (
	"fmt"
)

type operation func(args ...interface{}) interface{}

var operations map[string]operation = make(map[string]operation)

func opCustom(opName string, arg interface{}, data interface{}) (interface{}, error) {
	if op, ok := operations[opName]; ok {
		switch arg.(type) {
		case []interface{}:
			var err error
			argValues := make([]interface{}, len(arg.([]interface{})))
			for i, argExpr := range arg.([]interface{}) {
				argValues[i], err = applyInterfaces(argExpr, data)
				if err != nil {
					return nil, fmt.Errorf("error")
				}
			}
			return op(argValues...), nil
		default:
			args, err := applyInterfaces(arg, data)
			if err != nil {
				return nil, fmt.Errorf("error")
			}
			return (op(args)), nil
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
