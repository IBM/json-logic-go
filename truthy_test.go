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

import "testing"
import "github.com/stretchr/testify/assert"

func TestTruthy(t *testing.T) {
	var ok bool
	var err error

	ok, err = truthy(false)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy(true)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(0)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy(-1)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(1)
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy([]interface{}{})
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy([]interface{}{1, 2})
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy("")
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	ok, err = truthy("anything")
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy("0")
	assert.Equal(t, true, ok)
	assert.NoError(t, err)

	ok, err = truthy(nil)
	assert.Equal(t, false, ok)
	assert.NoError(t, err)

	// Unsupported type
	ok, err = truthy([]byte{45, 32})
	assert.Equal(t, false, ok)
	assert.Error(t, err)
}
