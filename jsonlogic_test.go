package jsonlogic

import (
	"testing"
)

func TestEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "==" : [1, 1] }`)
	result = Apply(f)

	if result != true {
		t.Error("expected true, got", result)
	}

	f = StringToInterface(`{ "==" : [1, 1.0] }`)
	result = Apply(f)

	if result != true {
		t.Error("expected true, got", result)
	}

	f = StringToInterface(`{ "==" : [1, 1.1] }`)
	result = Apply(f)

	if result != false {
		t.Error("expected false, got", result)
	}

	f = StringToInterface(`{ "==" : [1, 0] }`)
	result = Apply(f)

	if result != false {
		t.Error("expected false, got", result)
	}
}

func TestUnEqual(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`{ "!=" : [1, 0] }`)
	result = Apply(f)

	if result != true {
		t.Error("expected true, got", result)
	}

	f = StringToInterface(`{ "!=" : [1, 1] }`)
	result = Apply(f)

	if result != false {
		t.Error("expected false, got", result)
	}
}

func TestNonRule(t *testing.T) {
	var result interface{}
	var f interface{}

	f = StringToInterface(`true`)
	result = Apply(f)

	if result != true {
		t.Error("expected true, got", result)
	}

	f = StringToInterface(`false`)
	result = Apply(f)

	if result != false {
		t.Error("expected false, got", result)
	}

	f = StringToInterface(`17`)
	result = Apply(f)

	if result != 17.0 { // Note, it seems like all json numerics are float64
		t.Error("expected 17, got", result)
	}

	f = StringToInterface(`3.14`)
	result = Apply(f)

	if result != 3.14 {
		t.Error("expected 3.14, got", result)
	}

	f = StringToInterface("apple")
	result = Apply(f)

	if result != "apple" {
		t.Error("expected apple, got", result)
	}

	//TODO: I am skipping here a test for "null". I don't think golang can handle this corner case.

	//TODO: Arrays
	// f = StringToInterface(`["a", "b"]`)
	// result = Apply(f)

	// if reflect.TypeOf(result) == []string && testEq(result, []string{"a", "b"}) {
	// 	t.Error("expected [a,b], got", result)
	// }

}

// Helper methods
func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
