package jsonlogic

import "testing"

func TestEqual(t *testing.T) {
	var result bool
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
	var result bool
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
