package jsonlogic

import "testing"
import "github.com/stretchr/testify/assert"

func TestTruthy(t *testing.T) {
	var ok bool
	var err error

	ok, err = truthy(false)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy(true)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(0)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy(-1)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(1)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy([]interface{}{})
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy([]interface{}{1, 2})
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy("")
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy("anything")
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy("0")
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(nil)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	// Unsupported type
	ok, err = truthy([]byte{45, 32})
	assert.Equal(t, false, ok)
	assert.Error(t, err)
}
