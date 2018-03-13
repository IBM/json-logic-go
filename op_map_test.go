package jsonlogic

import "testing"
import "github.com/stretchr/testify/assert"

func TestHardcodedMap(t *testing.T) {
	var result interface{}

	result = Apply(`{"map":[[1,2,3,4,5],{"*":[{"var":""},2]}]}`)
	assert.Equal(t, []float64{2, 4, 6, 8, 10}, result)
}
