package jsonlogic

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Address struct {
	Country string
	State   string
}

type Person struct {
	FirstName string
	LastName  string
	Age       uint8
	Address   Address
}

func TestVarWithStruct(t *testing.T) {
	p := Person{
		FirstName: "John",
		LastName:  "Smith",
		Age:       22,
		Address: Address{
			Country: "US",
			State:   "MA",
		},
	}

	var rule interface{}

	err := json.Unmarshal([]byte(`{">": [{"var": "Age"}, 21]}`), &rule)

	assert.NoError(t, err)

	result, err := ApplyJSONInterfaces(rule, p)

	assert.NoError(t, err)

	assert.True(t, result.(bool))
}

func TestDataDrivenRule1(t *testing.T) {
	result, err := Apply(`{ "var": ["a", 1] }`)

	assert.NoError(t, err)

	assert.Equal(t, float64(1), result)
}

func TestDataDrivenRule2(t *testing.T) {
	result, err := Apply(`{"var":["b",2]}`, `{"a":1}`)

	assert.NoError(t, err)

	assert.Equal(t, float64(2), result)
}
