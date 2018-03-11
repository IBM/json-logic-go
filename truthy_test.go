package jsonlogic

import "testing"
import "github.com/stretchr/testify/assert"

func TestTruthy(t *testing.T) {
	assert.Equal(t, false, truthy(false))
	assert.Equal(t, true, truthy(true))
	assert.Equal(t, false, truthy(0))
	assert.Equal(t, true, truthy(-1))
	assert.Equal(t, true, truthy(1))
	assert.Equal(t, false, truthy([]interface{}{}))
	assert.Equal(t, true, truthy([]interface{}{1, 2}))
	assert.Equal(t, false, truthy(""))
	assert.Equal(t, true, truthy("anything"))
	assert.Equal(t, true, truthy("0"))
	assert.Equal(t, false, truthy(nil))
}
