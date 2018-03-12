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
}
