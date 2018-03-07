# json-logic-go

This is an attempt to port json-logic (http://jsonlogic.com/) to golang.

## API
- `Apply(rule string, data string)`: See json-logic docs for details on the rules syntax and the data syntax. Data is optional.

## Supported operators
- `var`
- `==`
- `===` (alias of `==`, which is not the case in the original jsonlogic)
- `!=`
- `!==` (alias of `!=`, which is not the case in the original jsonlogic)
- `and`

## Known issues
- Dot operation is not supported yet. You cannot use data that represents an object. 
- They also have some hybrid dot notation to access an object within an array or something like that. Not supported.
- json-logic officially considers that `1 == "1"`. Not supported here.
- json-logic makes a difference between `==` and `===`. Not supported here.
- The expected result of a `null` rule is `null`. I don't think it's true here.
- In javascript (and hence json-logic) `1 && 3 == 3`. Not supported here.
- json-logic considers `{"var":"1"}` and `{"var":1}` to be the same. Not supported here. Var names need to be either strings or numbers.
- Order of operations? I did not research nor test that the order of operations is always correct.
- Need to check if we obey the same rules of truthy (http://jsonlogic.com/truthy.html)
- Error handling? 

## Unimplemented operators
- `missing`
- `missing_some`
- `if`
- `!`
- `!!`
- `or`
- `>`
- `>=`
- `<`
- `<=`
- Between
- `max`
- `min`
- `+`
- `-`
- `*`
- `/`
- `%`
- `map`
- `reduce`
- `filter`
- `all`
- `none`
- `some`
- `merge`
- `in`
- `cat`
- `substr`
- `log`
- Custom Operators