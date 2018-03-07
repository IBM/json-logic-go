package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "==" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, result, true)

	f = StringToInterface(`{ "==" : [1, 1.0] }`)
	result = Apply(f)
	assert.Equal(t, result, true)

	f = StringToInterface(`{ "==" : [1, 1.1] }`)
	result = Apply(f)
	assert.Equal(t, result, false)

	f = StringToInterface(`{ "==" : [1, 0] }`)
	result = Apply(f)
	assert.Equal(t, result, false)

	// f = StringToInterface(`{ "==" : [1, "1"] }`)
	// result = Apply(f)

	// if result != true {
	// 	t.Error("expected true, got", result)
	// }
}

func TestUnEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "!=" : [1, 0] }`)
	result = Apply(f)
	assert.Equal(t, result, true)

	f = StringToInterface(`{ "!=" : [1, 1] }`)
	result = Apply(f)
	assert.Equal(t, result, false)
}

func TestNonRule(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`true`)
	result = Apply(f)
	assert.Equal(t, result, true)

	f = StringToInterface(`false`)
	result = Apply(f)
	assert.Equal(t, result, false)

	f = StringToInterface(`17`)
	result = Apply(f)
	assert.Equal(t, result, float64(17))

	f = StringToInterface(`3.14`)
	result = Apply(f)
	assert.Equal(t, result, 3.14)

	f = StringToInterface("apple")
	result = Apply(f)
	assert.Equal(t, result, "apple")

	//TODO: I am skipping here a test for "null". I don't think golang can handle this corner case.

	f = StringToInterface(`["a", "b"]`)
	result = Apply(f)
	var targetValue = []interface{}{"a", "b"} //DeepEqual only works for the same types
	assert.Equal(t, result, targetValue)

}
