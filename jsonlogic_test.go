/*
MIT License

Copyright (c) 2018 IBM

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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

// Tests as defined in http://jsonlogic.com/tests.json, (skipped unknown variable test)
func TestRemote(t *testing.T) {
	var testData interface{}
	err := getLocalJSON("tests.json", &testData)
	if err != nil {
		log.Fatal("Failed to load local tests, stop!")
	}

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
			actual, _ := ApplyJSONInterfaces(rule, data) //Ignoring error, I don't know which cases expect errors or not ...
			ok := assert.Equal(t, expected, actual, json)
			if ok {
				success++
			}
		default:
			//Skip comments
			total--
		}
	}

	log.Println(success, "success out of", total)

}

// Tests not covered by tests.json
func TestJSON(t *testing.T) {
	var result interface{}
	var err error

	// Invalid json rule
	result, err = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]`)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Invalid json data
	result, err = Apply(`{"if": [{"var": "a"}, "yes", "no"]}`, `{"a": "var": "a"}}`)
	assert.Error(t, err)
	assert.Nil(t, result)

	result, err = Apply(`{"if": [{"var": "a"}, "yes", "no"]}`, `[`)
	assert.Error(t, err)
	assert.Nil(t, result)

	// Unknown operator
	result, err = Apply(`{"if": [{"hOOUIGOFoyggf": "a"}, "yes", "no"]}`, `{"a": {"var": "a"}}`)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnSupportedType(t *testing.T) {
	result, err := opAll([]byte{12}, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opAll([]interface{}{nil}, 1)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opAnd(nil, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opAnd([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opSum(true, nil)
	assert.Equal(t, float64(0), result)

	result, err = opSum([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, float64(0), result)

	result, err = opCat([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opEqual(3, []byte{})
	assert.Equal(t, false, result)

	result, err = opEqual([]byte{}, []byte{})
	assert.Equal(t, false, result)

	result, err = opEqualStrict(3, []byte{})
	assert.Equal(t, false, result)

	result, err = opEqualStrict([]byte{}, []byte{})
	assert.Equal(t, false, result)

	result, err = opNotEqual(3, []byte{})
	assert.Equal(t, true, result)

	result, err = opNotEqual([]byte{}, []byte{})
	assert.Equal(t, true, result)

	result, err = opNotEqualStrict(3, []byte{})
	assert.Equal(t, true, result)

	result, err = opNotEqualStrict([]byte{}, []byte{})
	assert.Equal(t, true, result)

	result, err = opSmallerThan(1, []byte{})
	assert.Equal(t, false, result)

	result, err = opSmallerEqThan([]byte{}, []byte{})
	assert.Equal(t, true, result)

	result, err = opGreaterThan(1, []byte{})
	assert.Equal(t, true, result)

	result, err = opGreaterEqThan([]byte{}, nil)
	assert.Equal(t, true, result)

	result, err = opFilter(1, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opFilter([]byte{42}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opIf([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, []uint8([]byte{}), result)

	result, err = opIf(2, []interface{}{[]byte{}})
	assert.Equal(t, 2, result)

	result, err = opIn([]byte{}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opIn(1, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opLog([]byte{})
	assert.Equal(t, []uint8([]byte{}), result)

	result, err = opMap(nil, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opMap([]byte{}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opMap(1, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opMap([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opMissingSome([]interface{}{[]byte{}}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opMissingSome([]byte{}, []interface{}{[]byte{}})
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opNone(1, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opOr(1, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opOr([]byte{12}, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opReduce([]byte{12}, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opSome([]byte{12}, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opSome(1, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opSubstr(1, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = opSubstr([]interface{}{[]byte{45, 32}}, nil)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

}
func TestCompound(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,2]}]}]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"and":[{"==":[1,1]},{"and":[{"==":[1,1]},{"==":[2,1]}]}]}`)
	assert.Equal(t, false, result)

}
func TestVar(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"var": "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y"}`, `{"a":{"b":{"c":{"d":{"e":{"f":{"g":{"h":{"i":{"j":{"k":{"l":{"m":{"n":{"o":{"p":{"q":{"r":{"s":{"t":{"u":{"v":{"w":{"x":{"y":"z"}}}}}}}}}}}}}}}}}}}}}}}}}`)
	assert.Equal(t, "z", result)
	assert.NoError(t, err)

	result, err = Apply(`{"if": [{"var": "a"}, "yes", "no"]}`, `{"a": {"var": "a"}}`)
	assert.Equal(t, "yes", result)
	assert.NoError(t, err)

	result, err = Apply(`{"var": "a"}`)
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestLog(t *testing.T) {
	var result interface{}
	result, _ = Apply(`{"log":"log test"}`)
	assert.Equal(t, "log test", result)

	result, _ = Apply(`{"log":{"substr": [" log test_2 ", 1, -1]}}`)
	assert.Equal(t, "log test_2", result)
}
func TestMax(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"max": [{"var": "current"}, {"+":  [{"var":  "accumulator"}, 100] }]}`)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

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

	result, _ = Apply(`{"max": {"var": "a"}}`, `{"a": [1, 2, 3]}`)
	assert.Equal(t, 3.0, result)

	result, _ = Apply(`{"max": [{"var": "a"}, {"var": "b"}, {"var": "c"}]}`, `{"a": 1, "b": 2, "c": 3}`)
	assert.Equal(t, 3.0, result)

}

func TestMin(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"min": [{"var": "current"}, {"+":  [{"var":  "accumulator"}, 100] }]}`)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

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

	result, _ = Apply(`{"min": {"var": "a"}}`, `{"a": [1, 2, 3]}`)
	assert.Equal(t, 1.0, result)

	result, _ = Apply(`{"min": [{"var": "a"}, {"var": "b"}, {"var": "c"}]}`, `{"a": 1, "b": 2, "c": 3}`)
	assert.Equal(t, 1.0, result)

}

func TestMap(t *testing.T) {
	var result interface{}
	var err error
	result, err = Apply(`{"map":[{"var":"integers"}]}`, `{"integers":[]}`)
	assert.Equal(t, []interface{}{}, result)

	result, err = Apply(`{"map":[]}`)
	assert.Equal(t, []interface{}{}, result)

	result, err = Apply(`{"map":[{"var":"integers"}]}`, `{"integers":[1,2,null,4,5]}`)
	assert.Equal(t, []interface{}{nil, nil, nil, nil, nil}, result)
	assert.NoError(t, err)

	result, err = Apply(`{"map":[{"var":"integers"},{"*":[{"var":""},2]}]}`, `{"integers":[1,2,null,4,5]}`)
	assert.Equal(t, []interface{}{2.0, 4.0, nil, 8.0, 10.0}, result)
	assert.NoError(t, err)

	result, err = Apply(`{"map":[[1,2,3,4,5],{"*":[{"var":""},2]}]}`)
	assert.Equal(t, []interface{}{2.0, 4.0, 6.0, 8.0, 10.0}, result)
	assert.NoError(t, err)

	result, err = Apply(`{"map":[{"var": "a"},{"*":[{"var":""},2]}]}`, `{"a": [1,2,3,4,5]}`)
	assert.Equal(t, []interface{}{2.0, 4.0, 6.0, 8.0, 10.0}, result)
	assert.NoError(t, err)

	result, err = Apply(`{"map":[{"var": "a"},{"*":[{"var":""},2]}]}`)
	assert.Nil(t, result)
	assert.Error(t, err)

	result, _ = Apply(`{"map":[{"var": "b"},{"*":[{"var":"b"},2]}]}`, `{"a": [1,2,3,4,5]}`)
	assert.Nil(t, result)
	assert.Error(t, err)

	result, err = Apply(`{"map":[{"var": "a"},{"*":[{"var":"b"},2]}]}`, `{"a": [1,2,3,4,5]}`)
	assert.Nil(t, result)
	assert.Error(t, err)

	result, err = Apply(`{"map":[{"var": "a"},{"*":[{"var":"a"},2]}]}`, `{"a": [1,2,3,4,5]}`)
	assert.Nil(t, result)
	assert.Error(t, err)

}

func TestArithmetic(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{"/":[4,0]}`)
	assert.Equal(t, nil, result)
}
func TestIf(t *testing.T) {
	var result interface{}

	result, _ = Apply(`{ "if" : [] }`)
	assert.Equal(t, nil, result)

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

	result, _ = Apply(`{ "cat" : ["Hello, ",{"var":""}] }`, `"Dolly"`)
	assert.Equal(t, "Hello, Dolly", result)

	result, _ = Apply(`{ "cat" : [{"var":""}] }`, `"Dolly"`)
	assert.Equal(t, "Dolly", result)

	result, _ = Apply(`{ "cat" : {"var":""} }`, `Dolly`)
	assert.Equal(t, nil, result)

	result, _ = Apply(`{ "cat" : {"var":""} }`, `"Dolly"`)
	assert.Equal(t, "Dolly", result)

	result, _ = Apply(`{ "cat" : {"var": "a"} }`, `{"a": ["Hello, ","Dolly"]}`)
	assert.Equal(t, `Hello, Dolly`, result)
}

func TestReduce(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"reduce":[[50, 100, 150]]}`)
	assert.Equal(t, nil, result)

	result, err = Apply(`{"reduce":[[50, 100, 150],10]}`)
	assert.Equal(t, float64(10), result)
	assert.NoError(t, err)

	result, err = Apply(`{"reduce":[[50, 100, 150],{"max": [{"var": "current"}, {"+":  [{"var":  "accumulator"}, 100] }]}]}`)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

	result, err = Apply(`{"reduce":[[true, true, true],{"and": [{"var": "current"},{"var": "accumulator"}]}]}`)
	assert.Equal(t, nil, result)
	assert.Error(t, err)

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

	result, _ = Apply(`{"reduce":[
		{"var": "a"},
	   {"max": [{"var": "current"}, {"+":  [{"var":  "accumulator"}, 100] }]},
	   0
	]}`, `{"a": [50, 100, 150]}`)
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

	result, err = Apply(`{
		"reduce":[{"var":"desserts"},{"+":[{"var":"accumulator"},{"var":"current.wrong"}]},0]}`,
		`{"desserts":[{"name":"apple","qty":1},{"name":"brownie","qty":2},{"name":"cupcake","qty":3}]}`)
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestIn(t *testing.T) {
	var result interface{}
	result, _ = Apply(`{"in": ["Hello", {"var": "a"}]}`, `{"a": ["Hello", "World"]}`)
	assert.Equal(t, true, result)

	result, _ = Apply(`{"in": [5, {"var": "a"}]}`, `{"a": [5, 10]}`)
	assert.Equal(t, true, result)

}

func TestMerge(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"merge": [{"var": "a"}, {"var": "b"}]}`, `{"a": [1,2,3], "b": [4,5,6]}`)
	assert.Equal(t, []interface{}{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}, result)
	assert.NoError(t, err)

	result, err = Apply(`{"merge": [{"var":"integers"}, [1]]}`)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestFizzBuzz(t *testing.T) {
	var result interface{}
	var err error

	result, err = Apply(`{"map":[[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30],{"if":[{"==":[{"%":[{"var":""},15]},0]},"fizzbuzz",{"==":[{"%":[{"var":""},3]},0]},"fizz",{"==":[{"%":[{"var":""},5]},0]},"buzz",{"var":""}]}]}`)
	expected, _ := stringToInterface(`[1,2,"fizz",4,"buzz","fizz",7,8,"fizz","buzz",11,"fizz",13,14,"fizzbuzz",16,17,"fizz",19,"buzz","fizz",22,23,"fizz","buzz",26,"fizz",28,29,"fizzbuzz"]`)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	result, err = Apply(`{"map":[{"var":"a"},{"if":[{"==":[{"%":[{"var":""},15]},0]},"fizzbuzz",{"==":[{"%":[{"var":""},3]},0]},"fizz",{"==":[{"%":[{"var":""},5]},0]},"buzz",{"var":""}]}]}`,
		`{"a":[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30]}`)
	expected, _ = stringToInterface(`[1,2,"fizz",4,"buzz","fizz",7,8,"fizz","buzz",11,"fizz",13,14,"fizzbuzz",16,17,"fizz",19,"buzz","fizz",22,23,"fizz","buzz",26,"fizz",28,29,"fizzbuzz"]`)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

// Helper functions
func getJSON(url string, target interface{}) error {
	var client = http.Client{Timeout: 100 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	if res.StatusCode != 200 {
		err = fmt.Errorf("Request returned status code %v", res.StatusCode)
		log.Println(err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, &target)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func getLocalJSON(filename string, target interface{}) error {
	fileBody, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(fileBody, &target)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
