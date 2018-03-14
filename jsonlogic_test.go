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
			ok := assert.Equal(t, expected, applyInterfaces(rule, data), json)
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

	result = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]}`)
	assert.Equal(t, true, result)

	result = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,1]}]}]}`)
	assert.Equal(t, false, result)

	rule := map[string]interface{}{
		"and": []interface{}{
			map[string]interface{}{
				"==": []interface{}{1, 1},
			},
			map[string]interface{}{
				"and": []interface{}{
					map[string]interface{}{
						"==": []interface{}{1, 1},
					},
					map[string]interface{}{
						"==": []interface{}{2, 1},
					},
				},
			},
		},
	}
	result = Apply(rule)
	assert.Equal(t, false, result)

}

func TestVar(t *testing.T) {
	var result interface{}

	result = Apply(`{"var": "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y"}`, `{"a":{"b":{"c":{"d":{"e":{"f":{"g":{"h":{"i":{"j":{"k":{"l":{"m":{"n":{"o":{"p":{"q":{"r":{"s":{"t":{"u":{"v":{"w":{"x":{"y":"z"}}}}}}}}}}}}}}}}}}}}}}}}}`)
	assert.Equal(t, "z", result)

	result = Apply(`{"if": [{"var": "a"}, "yes", "no"]}`, `{"a": {"var": "a"}}`)
	assert.Equal(t, "yes", result)
}

func TestMax(t *testing.T) {
	var result interface{}

	result = Apply(`{ "max" : [1,2] }`)
	assert.Equal(t, float64(2), result)

	result = Apply(`{"max":[1,2,3]}`)
	assert.Equal(t, float64(3), result)

	result = Apply(`{"max":[]}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"max":["1"]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"max":["notnumber"]}`)
	assert.Equal(t, nil, result)

}

func TestMin(t *testing.T) {
	var result interface{}

	result = Apply(`{ "min" : [1,2] }`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":[1,2,3]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":[]}`)
	assert.Equal(t, nil, result)

	result = Apply(`{"min":["1"]}`)
	assert.Equal(t, float64(1), result)

	result = Apply(`{"min":["notnumber"]}`)
	assert.Equal(t, nil, result)

}

func TestMap(t *testing.T) {
	var result interface{}

	result = Apply(`{"map":[[1,2,3,4,5],{"*":[{"var":""},2]}]}`)
	assert.Equal(t, []interface{}{2.0, 4.0, 6.0, 8.0, 10.0}, result)
}

func TestIf(t *testing.T) {
	var result interface{}

	result = Apply(`{ "if" : [true, true, false] }`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : [false, true, false] }`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [true, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [{"==": [1, 1]}, "yes", "no"] }`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [{"==": [1, 2]}, "yes", "no"] }`)
	assert.Equal(t, "no", result)

	result = Apply(`{ "if" : []}`)
	assert.Equal(t, nil, result)

	result = Apply(`{ "if" : null}`)
	assert.Equal(t, nil, result)

	result = Apply(`{ "if" : true}`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : false}`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}]}`)
	assert.Equal(t, true, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}]}`)
	assert.Equal(t, false, result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 1]}, "yes", "no"]}`)
	assert.Equal(t, "yes", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", "no"]}`)
	assert.Equal(t, "no", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", true, "yes-2"]}`)
	assert.Equal(t, "yes-2", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, 2]}, "yes", {"!=": [1, 1]}, "yes-2", "no-2"]}`)
	assert.Equal(t, "no-2", result)

	result = Apply(`{ "if" : [false, "yes", {"==": [1, {"var": "a"}]}]}`, `{"a": 1}`)
	assert.Equal(t, true, result)

	result = Apply(`{"?:":[true,1,2]}`)
	assert.Equal(t, float64(1), result)
}

func TestCat(t *testing.T) {
	var result interface{}

	result = Apply(`{ "cat" : ["Hello, ",{"var":""}] }`, `Dolly`)
	assert.Equal(t, `Hello, Dolly`, result)

	result = Apply(`{ "cat" : [{"var":""}] }`, `Dolly`)
	assert.Equal(t, `Dolly`, result)

	result = Apply(`{ "cat" : {"var":""} }`, `Dolly`)
	assert.Equal(t, `Dolly`, result)
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
