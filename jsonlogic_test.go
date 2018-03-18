package jsonlogic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Tests as defined in http://jsonlogic.com/tests.json
func TestRemote(t *testing.T) {
	var testData interface{}
	getJSON("http://jsonlogic.com/tests.json", &testData)

	testDataArray := testData.([]interface{})

	var rule, data, expected interface{}
	total := len(testDataArray)
	success := 0

	for _, test := range testDataArray {
		switch test.(type) {
		case []interface{}:
			rule = test.([]interface{})[0]
			data = test.([]interface{})[1]
			expected = test.([]interface{})[2]
			b := new(bytes.Buffer)
			enc := json.NewEncoder(b)
			enc.SetEscapeHTML(false)
			enc.Encode(test)
			json := b.String()
			actual, err := applyInterfaces(rule, data)
			if err != nil {
				//TODO: check errors
			}
			ok := assert.Equal(t, expected, actual, json)
			if ok {
				success++
			}
		default:
			//Skip comments
			total--
		}
	}

	defer fmt.Println(success, "success out of", total)

}

// Tests not covered by tests.json
func TestCompound(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,1]}]}]}`)
	assert.Equal(t, false, result)

}
func TestVar(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{"var": "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y"}`, `{"a":{"b":{"c":{"d":{"e":{"f":{"g":{"h":{"i":{"j":{"k":{"l":{"m":{"n":{"o":{"p":{"q":{"r":{"s":{"t":{"u":{"v":{"w":{"x":{"y":"z"}}}}}}}}}}}}}}}}}}}}}}}}}`)
	assert.Equal(t, "z", result)

	result, _ = Apply(`{"if": [{"var": "a"}, "yes", "no"]}`, `{"a": {"var": "a"}}`)
	assert.Equal(t, "yes", result)
}

func TestMax(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{ "max" : [1,2] }`)
	assert.Equal(t, float64(2), result)

	result, _ = Apply(`{"max":[1,2,3]}`)
	assert.Equal(t, float64(3), result)

	result, _ = Apply(`{"max":[]}`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{"max":["1"]}`)
	assert.Equal(t, float64(1), result)

	result, _ = Apply(`{"max":["notnumber"]}`)
	assert.Equal(t, nil, result)

}

func TestMin(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{ "min" : [1,2] }`)
	assert.Equal(t, float64(1), result)

	result, _ = Apply(`{"min":[1,2,3]}`)
	assert.Equal(t, float64(1), result)

	result, _ = Apply(`{"min":[]}`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{"min":["1"]}`)
	assert.Equal(t, float64(1), result)

	result, _ = Apply(`{"min":["notnumber"]}`)
	assert.Equal(t, nil, result)

}

func TestMap(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{"map":[[1,2,3,4,5],{"*":[{"var":""},2]}]}`)
	assert.Equal(t, []interface{}{2.0, 4.0, 6.0, 8.0, 10.0}, result)
}

func TestIf(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{ "if" : [true, true, false] }`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{ "if" : [false, true, false] }`)
	assert.Equal(t, false, result)

	result, _ = Apply(`{ "if" : [true, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result, _ = Apply(`{ "if" : [{"==": [1, 1]}, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result, _ = Apply(`{ "if" : [{"==": [1, 2]}, "yes", "no"] }`)
	assert.Equal(t, "no", result)

	result, _ = Apply(`{ "if" : []}`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{ "if" : null}`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{ "if" : true}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{ "if" : false}`)
	assert.Equal(t, false, result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}]}`)
	assert.Equal(t, false, result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}, "yes", "no"]}`)
	assert.Equal(t, "yes", result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", "no"]}`)
	assert.Equal(t, "no", result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", true, "yes-2"]}`)
	assert.Equal(t, "yes-2", result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", {"!=": [1, 1]}, "yes-2", "no-2"]}`)
	assert.Equal(t, "no-2", result)

	result, _ = Apply(`{ "if" : [false, "yes", {"==": [1, {"var": "a"}]}]}`, `{"a": 1}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"?:":[true,1,2]}`)
	assert.Equal(t, float64(1), result)
}

func TestCat(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{ "cat" : ["Hello, ",{"var":""}] }`, `Dolly`)
	assert.Equal(t, `Hello, Dolly`, result)

	result, _ = Apply(`{ "cat" : [{"var":""}] }`, `Dolly`)
	assert.Equal(t, `Dolly`, result)

	result, _ = Apply(`{ "cat" : {"var":""} }`, `Dolly`)
	assert.Equal(t, `Dolly`, result)
}

func TestReduce(t *testing.T) {
	var result interface{}
	result, _ = Apply(`{"reduce":[[true, true, true],{"and": [{"var": "current"},{"var": "accumulator"}]},true]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"reduce":[[true, true, false],{"and": [{"var": "current"},{"var": "accumulator"}]},true]}`)
	assert.Equal(t, false, result)

	result, _ = Apply(`{"reduce":[
		[50, 100, 150],
	   {"max": [{"var": "current"}, {"+":  [{"var":  "accumulator"}, 100] }]},
	   0
	]}`)
	assert.Equal(t, float64(300), result)

	// From the jsonlogic doc:
	// Note, that inside the logic being used to reduce, var operations only have access to an object like:
	// {
	// 	"current" : // this element of the array,
	// 	"accumulator" : // progress so far, or the initial value
	// }
	//
	// This rule should evaluate to nil because var operation does not comply with this constraint
	result, _ = Apply(`{
		"reduce": [
			[1,2,3],
			{"+": [{"var": "a"}, {"var": "b"}]},
			10
		]
		}`,
		`{
		"a": 100, 
		"b": 1000
	}`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{
		"reduce": [
			[1,2,3],
			{"+": [{"var": "accumulator"}, {"var": "accumulator"}]},
			10
		]
		}`,
		`{
		"a": 100,
		"b": 1000
	}`)
	assert.Equal(t, float64(80), result)

	result, _ = Apply(`{
		"reduce":[{"var":"desserts"},{"+":[{"var":"accumulator"},{"var":"current.wrong"}]},0]}`,
		`{"desserts":[{"name":"apple","qty":1},{"name":"brownie","qty":2},{"name":"cupcake","qty":3}]}`)
	assert.Equal(t, nil, result)
}

func TestIn(t *testing.T) {
	var result interface{}
	result, _ = Apply(`{"in": ["Hello", {"var": "a"}]}`, `{"a": ["Hello", "World"]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"in": [5, {"var": "a"}]}`, `{"a": [5, 10]}`)
	assert.Equal(t, true, result)
}

// Helper function
func getJSON(url string, target interface{}) {
	var client = http.Client{Timeout: 100 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &target)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}
