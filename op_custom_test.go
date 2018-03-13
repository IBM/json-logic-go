package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func myOp(args ...interface{}) interface{} {
	return "Hi " + args[0].(string) + "!"
}

func add(args ...interface{}) interface{} {
	x, y := float64(args[0].(float64)), float64(args[1].(float64))

	return x + y
}

func TestAddOperation(t *testing.T) {
	err := AddOperation("myOp", myOp)
	assert.NoError(t, err)

	result := Apply(`{"myOp": "jsonlogic"}`)
	assert.Equal(t, "Hi jsonlogic!", result)

	AddOperation("add", add)
	result = Apply(`{"add": [1, 2]}`)
	assert.Equal(t, float64(3), result)
}