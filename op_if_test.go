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

	// result = Apply(`{ "if" : [false, "yes", {">": [2, 1]}]}`)
	// assert.Equal(t, true, result)

	// result = Apply(`{ "if" : []}`)
	// assert.Equal(t, nil, result)

	// result = Apply(`{ "if" : null}`)
	// assert.Equal(t, nil, result)

	// result = Apply(`{ "if" : true}`)
	// assert.Equal(t, true, result)

	// result = Apply(`{ "if" : false}`)
	// assert.Equal(t, false, result)

}
