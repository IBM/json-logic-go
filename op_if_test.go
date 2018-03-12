package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	var result interface{}

	result = Apply(`{ "if" : [true, true, false] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : [false, true, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [true, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [{"==": [1, 1]}, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [{"==": [1, 2]}, "yes", "no"] }`)
	assert.Equal(t, "no", result)

	result = Apply(`{ "if" : []}`)
	assert.Equal(t, nil, result)

	result = Apply(`{ "if" : null}`)
	assert.Equal(t, nil, result)

	result = Apply(`{ "if" : true}`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : false}`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}]}`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}]}`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}, "yes", "no"]}`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", "no"]}`)
	assert.Equal(t, "no", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", true, "yes-2"]}`)
	assert.Equal(t, "yes-2", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", {"!=": [1, 1]}, "yes-2", "no-2"]}`)
	assert.Equal(t, "no-2", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, {"var": "a"}]}]}`, `{"a": 1}`)
	assert.Equal(t, true, result)

	result = Apply(`{"?:":[true,1,2]}`)
	assert.Equal(t, float64(1), result)
}
