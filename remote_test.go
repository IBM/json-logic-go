package jsonlogic

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
func TestAll(t *testing.T) {
	var testData interface{}
	getJSON("http://jsonlogic.com/tests.json", &testData)

	testDataArray := testData.([]interface{})

	var rule, data, expected interface{}

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
			assert.Equal(t, expected, applyInterfaces(rule, data), json)
		default:
			//Skip comments
		}
	}

}
