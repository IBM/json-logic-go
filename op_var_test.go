package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeep(t *testing.T) {
	var result interface{}

	result = Apply(`{"var": "a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y"}`, `{"a":{"b":{"c":{"d":{"e":{"f":{"g":{"h":{"i":{"j":{"k":{"l":{"m":{"n":{"o":{"p":{"q":{"r":{"s":{"t":{"u":{"v":{"w":{"x":{"y":"z"}}}}}}}}}}}}}}}}}}}}}}}}}`)
	assert.Equal(t, "z", result)
}
