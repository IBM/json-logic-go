[![Build Status](https://travis.ibm.com/NATHANH/json-logic-go.svg?token=DxTFrpuyhWKJNXaj5cp2&branch=master)](https://travis.ibm.com/NATHANH/json-logic-go)
# json-logic-go

This parser accepts [JsonLogic](http://jsonlogic.com/) rules and executes them in Go.

The JsonLogic format is designed to allow you to share rules (logic) between front-end and back-end code (regardless of language difference), even to store logic along with a record in a database. JsonLogic is documented extensively at JsonLogic.com, including examples of every [supported operation](http://jsonlogic.com/operations.html) and a place to [try out rules in your browser](http://jsonlogic.com/play.html).

# API
There is currently only one exposed API to this module:

## Apply
`func Apply(inputs ...string) (interface{}, error)`

The first input is the json rule, as a string:

``Apply(`{ "==" : [1, 1] } `);``

The second input, optional, is the data source:

``Apply(`{ "var" : ["a"] }`, `{ a : 1, b: 2 }`);``

`Apply()` uses `json.Unmarshal` to convert the string inputs to generic `interface{}`.

The return value is of type `interface{}` because it can be anything such as a string, boolean, map, array, etc. The function may also return an error.

## ApplyJSONInterfaces
`func ApplyJSONInterfaces(inputs ...interface{}) (interface{}, error)`

This has the same behavior as `Apply()`, but expects inputs already unmarshalled. This should not be used with your own custom types, but rather ONLY with unmarshalled json.

## AddOperation
`func AddOperation(name string, implementation Operation) error`

This allows you to extend the range of operations provided by jsonlogic.  
An `Operation` has should have the following signature: 
`func(args ...interface{}) (interface{}, error)`

For example:

```
func add(args ...interface{}) (interface{}, error) {
	x, y := float64(args[0].(float64)), float64(args[1].(float64))

	return (x + y), nil
}

jsonlogic.AddOperation("add", add)
result, _ = Apply(`{"add": [1, 2]}`)
```

