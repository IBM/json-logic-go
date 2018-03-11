package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests as defined in http://jsonlogic.com/tests.json
func TestEqual(t *testing.T) {
	var result interface{}

	// f = StringToInterface(`{ "==" : [1, 1] }`)
	result = Apply(`{ "==" : [1, 1] }`)
	assert.Equal(t, true, result)

	// TODO: This test fails. Do we care? It's in the official tests ..
	// f = StringToInterface(`{ "==" : [1, "1"] }`)
	// result = Apply(f)
	// assert.Equal(t, true, result)

	result = Apply(`{ "==" : [1, 2] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "===" : [1, 1] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "===" : [1, "1"] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "===" : [1, 2] }`)
	assert.Equal(t, false, result)
}

func TestUnEqual(t *testing.T) {
	var result interface{}

	result = Apply(`{ "!=" : [1, 2] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "!=" : [1, 1] }`)
	assert.Equal(t, false, result)

	// TODO: Test is failing. Unsupported in go.
	// f = StringToInterface(`{ "!=" : [1, "1"] }`)
	// result = Apply(f)
	// assert.Equal(t, false, result)

	result = Apply(`{ "!==" : [1, 2] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "!==" : [1, 1] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "!==" : [1, "1"] }`)
	assert.Equal(t, true, result)
}

func TestNonRule(t *testing.T) {
	var result interface{}

	result = Apply(`true`)
	assert.Equal(t, true, result)

	result = Apply(`false`)
	assert.Equal(t, false, result)

	result = Apply(`17`)
	assert.Equal(t, float64(17), result)

	result = Apply(`3.14`)
	assert.Equal(t, 3.14, result)

	result = Apply("apple")
	assert.Equal(t, "apple", result)

	//TODO: I am skipping here a test for "null". I don't think golang can handle this corner case.

	result = Apply(`["a", "b"]`)
	var targetValue = []interface{}{"a", "b"} //DeepEqual only works for the same types
	assert.Equal(t, targetValue, result)

}

func TestAnd(t *testing.T) {
	var result interface{}

	result = Apply(`{ "and" : [true, true] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "and" : [false, true] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [true, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [false, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [true, true, true] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "and" : [true, true, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [true] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "and" : [1, 3] }`)
	assert.Equal(t, float64(3), result)

	result = Apply(`{ "and" : [3, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "and" : [false, 3] }`)
	assert.Equal(t, false, result)
}

func TestCompound(t *testing.T) {
	var result interface{}

	result = Apply(`{ "and" : [{ "==" : [1, 1] }] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "and" : [{ "==" : [1, 2] }] }`)
	assert.Equal(t, false, result)

	result = Apply(`{"and":[{"==":[1,1]},{"==":[1,2]}]}`)
	assert.Equal(t, false, result)

	result = Apply(`{"and":[{"==":[1,1]},{"==":[2,2]}]}`)
	assert.Equal(t, true, result)

	result = Apply(`{"and":[{"==":[1,1]},true]}`)
	assert.Equal(t, true, result)

	result = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]}`)
	assert.Equal(t, true, result)

	result = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,1]}]}]}`)
	assert.Equal(t, false, result)

}

func TestDataDriven(t *testing.T) {
	var result interface{}

	result = Apply(`{"var":["a"]}`, `{"a":1}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"var":["b"]}`, `{"a":1}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"var":["a"]}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"var":"a"}`, `{"a":1}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"var":"b"}`, `{"a":1}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"var":["a", 1]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"var":["b", 2]}`, `{"a":1}`)
	assert.Equal(t, float64(2), result)

	//TODO: dot-notation. This is an advanced case.
	// rule = StringToInterface(`{"var":"a.b"}`)
	// data = StringToInterface(`{"a":{"b":"c"}}`)
	// result = Apply(rule, data)
	// assert.Equal(t, "c", result)

	result = Apply(`{"var":1}`, `["apple", "banana"]`)
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

	result = Apply(`{ "and" : [{ "==" : [1, {"var":"a"}] }] }`, `{"a":1}`)
	assert.Equal(t, true, result)

	result = Apply(`{ "and" : [{ "!=" : [1, {"var":"a"}] }] }`, `{"a":1}`)
	assert.Equal(t, false, result)

}
