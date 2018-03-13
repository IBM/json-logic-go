# json-logic-go

This parser accepts [JsonLogic](http://jsonlogic.com/) rules and executes them in Go.

The JsonLogic format is designed to allow you to share rules (logic) between front-end and back-end code (regardless of language difference), even to store logic along with a record in a database. JsonLogic is documented extensively at JsonLogic.com, including examples of every [supported operation](http://jsonlogic.com/operations.html) and a place to [try out rules in your browser](http://jsonlogic.com/play.html).

# API
There is currently only one exposed API to this module:

## Apply
`func Apply(inputs ...string) interface{}`

The first input is the json rule, as a string:

``Apply(`{ "==" : [1, 1] } `);``

The second input, optional, is the data source:

``jsonLogic.apply(`{ "var" : ["a"] }`, `{ a : 1, b: 2 }`);``

The return value is of type `interface{}` because it can be anything such as a string, boolean, map, array, etc