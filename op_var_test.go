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
