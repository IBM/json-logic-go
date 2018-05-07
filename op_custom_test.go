package jsonlogic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func myOp(args ...interface{}) (interface{}, error) {
	return "Hi " + args[0].(string) + "!", nil
}

func add(args ...interface{}) (interface{}, error) {
	x, y := float64(args[0].(float64)), float64(args[1].(float64))

	return (x + y), nil
}

func getFirstWidget(args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, errors.New("Missing argument")
	}

	var widgets map[string]interface{}

	widgets, ok := args[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("Not a map")
	}

	widgetsArray := widgets["Widget"].([]interface{})

	return widgetsArray[0], nil

}

func TestAddOperation(t *testing.T) {
	err := AddOperation("myOp", myOp)
	assert.NoError(t, err)

	result, _ := Apply(`{"myOp": "jsonlogic"}`)
	assert.Equal(t, "Hi jsonlogic!", result)

	AddOperation("add", add)
	result, _ = Apply(`{"add": [1, 2]}`)
	assert.Equal(t, float64(3), result)

	result, _ = Apply(`{"add": [{"if": [true, -1, 1]}, 2]}`)
	assert.Equal(t, float64(1), result)

}

func TestGetWidgetOperation(t *testing.T) {
	err := AddOperation("getFirstWidget", getFirstWidget)
	assert.NoError(t, err)

	result, _ := Apply(`{"getFirstWidget": {"var": ""}}`, `{"Widget": ["hello", "world"]}`)
	assert.Equal(t, "hello", result)
}
