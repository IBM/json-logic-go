/*
MIT License

Copyright (c) 2018 IBM

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package jsonlogic

import (
	"fmt"
)

// Operation defines an operation with function signature:
// `func(args ...interface{}) (interface{}, error)`
// The first argument is the value(s) passed to the operation, the second argument is an optional data array.
type Operation func(args ...interface{}) (interface{}, error)

var operations = make(map[string]Operation)

func opCustom(opName string, arg interface{}, data interface{}) (interface{}, error) {
	if op, ok := operations[opName]; ok {
		switch arg.(type) {
		case []interface{}:
			var err error
			argValues := make([]interface{}, len(arg.([]interface{})))
			for i, argExpr := range arg.([]interface{}) {
				argValues[i], err = ApplyJSONInterfaces(argExpr, data)
				if err != nil {
					return nil, err
				}
			}
			return op(argValues...)
		default:
			// This extra level of recursion was giving troubles when the retrieve data is a map itself...
			// Commented out for now.
			// args, err := ApplyJSONInterfaces(arg, data)
			// if err != nil {
			// 	return nil, err
			// }
			// return op(args)
			return op(arg)
		}

	}

	return nil, fmt.Errorf("Unknown uperation: %s", opName)
}

// AddOperation allows you to add a custom operation that will run a Go function.
// The `implementation` must be an `Operation`. See type definition.
func AddOperation(name string, implementation Operation) error {
	if _, ok := operations[name]; ok {
		return fmt.Errorf("Operation exists: %s", name)
	}

	operations[name] = implementation

	return nil
}
