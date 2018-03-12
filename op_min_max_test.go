package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	var result interface{}

	result = Apply(`{ "max" : [1,2] }`)
	assert.Equal(t, float64(2), result)

	result = Apply(`{"max":[1,2,3]}`)
	assert.Equal(t, float64(3), result)

	result = Apply(`{"max":[]}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"max":["1"]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"max":["notnumber"]}`)
	assert.Equal(t, nil, result)

}

func TestMin(t *testing.T) {
	var result interface{}

	result = Apply(`{ "min" : [1,2] }`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":[1,2,3]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":[]}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"min":["1"]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":["notnumber"]}`)
	assert.Equal(t, nil, result)

}
