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
	Hobbies   []string
}

func createPerson() Person {
	return Person{
		FirstName: "John",
		LastName:  "Smith",
		Age:       22,
		Address: Address{
			Country: "US",
			State:   "MA",
		},
		Hobbies: []string{"Sports", "Movies", "History"},
	}
}

func TestVarWithStruct(t *testing.T) {
	var rule interface{}

	err := json.Unmarshal([]byte(`{">": [{"var": "Age"}, 21]}`), &rule)

	assert.NoError(t, err)

	p := createPerson()

	result, err := ApplyJSONInterfaces(rule, p)

	assert.NoError(t, err)

	assert.True(t, result.(bool))
}

func TestVarWithStructAndArray(t *testing.T) {
	var rule interface{}

	err := json.Unmarshal([]byte(`{"in": ["Movies", {"var": "Hobbies"}]}`), &rule)

	assert.NoError(t, err)

	p := createPerson()

	result, err := ApplyJSONInterfaces(rule, p)

	assert.NoError(t, err)

	assert.True(t, result.(bool))
}

type Product struct {
	Name   string
	Price  float32
	CatNum string
}

type Catalog struct {
	Title    string
	Products []Product
}

func createCatalog() Catalog {
	return Catalog{
		Title: "Gizmos",
		Products: []Product{
			Product{
				Name:   "Thingie",
				Price:  1.95,
				CatNum: "1001",
			},
			Product{
				Name:   "Gadget",
				Price:  20.5,
				CatNum: "1010",
			},
			Product{
				Name:   "Device",
				Price:  100.0,
				CatNum: "1100",
			},
		},
	}
}
func TestVarWithStructAndArrayOfStructs(t *testing.T) {
	var rule interface{}

	err := json.Unmarshal([]byte(`{">": [{"var": "catalog.Products.2.Price"}, 1.94]}`), &rule)

	assert.NoError(t, err)

	c := createCatalog()

	result, err := ApplyJSONInterfaces(
		rule,
		map[string]interface{}{
			"catalog": c,
		})

	assert.NoError(t, err)

	assert.True(t, result.(bool))
}

func TestVarFetchStructFromArrayField(t *testing.T) {
	var rule interface{}

	err := json.Unmarshal([]byte(`{"var": "catalog.Products.2"}`), &rule)

	assert.NoError(t, err)

	c := createCatalog()

	result, err := ApplyJSONInterfaces(
		rule,
		map[string]interface{}{
			"catalog": c,
		})

	assert.NoError(t, err)

	product, ok := result.(Product)

	assert.True(t, ok, "Result must be an instance of Product")

	assert.True(t, product.CatNum == "1100", "Cat number does not match")
}
