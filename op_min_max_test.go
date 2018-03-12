package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	var result interface{}

	result = Apply(`{ "max" : [1,2] }`)
	assert.Equal(t, 1, result)
}
