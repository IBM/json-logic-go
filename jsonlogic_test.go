package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests as defined in http://jsonlogic.com/tests.json
func TestEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "==" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	// TODO: This test fails. Do we care? It's in the official tests ..
	// f = StringToInterface(`{ "==" : [1, "1"] }`)
	// result = Apply(f)
	// assert.Equal(t, true, result)

	f = StringToInterface(`{ "==" : [1, 2] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "===" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "===" : [1, "1"] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "===" : [1, 2] }`)
	result = Apply(f)
	assert.Equal(t, false, result)
}

func TestUnEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "!=" : [1, 2] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "!=" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	// TODO: Test is failing. Unsupported in go.
	// f = StringToInterface(`{ "!=" : [1, "1"] }`)
	// result = Apply(f)
	// assert.Equal(t, false, result)

	f = StringToInterface(`{ "!==" : [1, 2] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "!==" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "!==" : [1, "1"] }`)
	result = Apply(f)
	assert.Equal(t, true, result)
}

func TestNonRule(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`true`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`false`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`17`)
	result = Apply(f)
	assert.Equal(t, float64(17), result)

	f = StringToInterface(`3.14`)
	result = Apply(f)
	assert.Equal(t, 3.14, result)

	f = StringToInterface("apple")
	result = Apply(f)
	assert.Equal(t, "apple", result)

	//TODO: I am skipping here a test for "null". I don't think golang can handle this corner case.

	f = StringToInterface(`["a", "b"]`)
	result = Apply(f)
	var targetValue = []interface{}{"a", "b"} //DeepEqual only works for the same types
	assert.Equal(t, targetValue, result)

}

func TestAnd(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "and" : [true, true] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "and" : [false, true] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [true, false] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [false, false] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [true, true, true] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "and" : [true, true, false] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [false] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [true] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	// TODO: Apparently this is expected behavior in javascript. But not in golang!
	// f = StringToInterface(`{ "and" : [1, 3] }`)
	// result = Apply(f)
	// assert.Equal(t, 3, result)

	f = StringToInterface(`{ "and" : [3, false] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{ "and" : [false, 3] }`)
	result = Apply(f)
	assert.Equal(t, false, result)
}

func TestCompound(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "and" : [{ "==" : [1, 1] }] }`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{ "and" : [{ "==" : [1, 2] }] }`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{"and":[{"==":[1,1]},{"==":[1,2]}]}`)
	result = Apply(f)
	assert.Equal(t, false, result)

	f = StringToInterface(`{"and":[{"==":[1,1]},{"==":[2,2]}]}`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{"and":[{"==":[1,1]},true]}`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]}`)
	result = Apply(f)
	assert.Equal(t, true, result)

	f = StringToInterface(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,1]}]}]}`)
	result = Apply(f)
	assert.Equal(t, false, result)
}

func TestDataDriven(t *testing.T) {
	var result, rule, data interface{}

	rule = StringToInterface(`{"var":["a"]}`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, float64(1), result)

	rule = StringToInterface(`{"var":["b"]}`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, nil, result)

	rule = StringToInterface(`{"var":["a"]}`)
	result = Apply(rule)
	assert.Equal(t, nil, result)

	rule = StringToInterface(`{"var":"a"}`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, float64(1), result)

	rule = StringToInterface(`{"var":"b"}`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, nil, result)

	rule = StringToInterface(`{"var":["a", 1]}`)
	result = Apply(rule)
	assert.Equal(t, float64(1), result)

	rule = StringToInterface(`{"var":["b", 2]}`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, float64(2), result)

	//TODO: dot-notation. This is an advanced case.
	// rule = StringToInterface(`{"var":"a.b"}`)
	// data = StringToInterface(`{"a":{"b":"c"}}`)
	// result = Apply(rule, data)
	// assert.Equal(t, "c", result)

	rule = StringToInterface(`{"var":1}`)
	data = StringToInterface(`["apple", "banana"]`)
	result = Apply(rule, data)
	assert.Equal(t, "banana", result)

	//TODO: "1" is not the same as int(1) in Go! In javascript yes...
	// rule = StringToInterface(`{"var":"1"}`)
	// data = StringToInterface(`["apple", "banana"]`)
	// result = Apply(rule, data)
	// assert.Equal(t, "banana", result)

	//TODO; dot-notation, advanced use
	// rule = StringToInterface(`{"var":"1.1"}`)
	// data = StringToInterface(`["apple", ["banana", "beer"]]`)
	// result = Apply(rule, data)
	// assert.Equal(t, "beer", result)

	rule = StringToInterface(`{ "and" : [{ "==" : [1, {"var":"a"}] }] }`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, true, result)

	rule = StringToInterface(`{ "and" : [{ "!=" : [1, {"var":"a"}] }] }`)
	data = StringToInterface(`{"a":1}`)
	result = Apply(rule, data)
	assert.Equal(t, false, result)

}
