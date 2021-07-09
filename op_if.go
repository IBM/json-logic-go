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

// The if statement typically takes 3 arguments: a condition (if), what to do if it’s true (then), and what to do if it’s false (else), like:
// {"if" : [ true, "yes", "no" ]}
//
// If can also take more than 3 arguments, and will pair up arguments like if/then elseif/then elseif/then else. Like:
// {"if" : [
// 	{"<": [{"var":"temp"}, 0] }, "freezing",
// 	{"<": [{"var":"temp"}, 100] }, "liquid",
// 	"gas"
// ]}

func opIf(value interface{}, data interface{}) (interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch value.(type) {
	case []interface{}: // An array of values
		valuearray := value.([]interface{})

		if len(valuearray) == 0 {
			return nil, nil
		}

		condition, err := ApplyJSONInterfaces(valuearray[0], data)
		if err != nil {
			return nil, err
		}

		if len(valuearray) == 1 {
			return condition, nil
		}

		ok, err := truthy(condition)
		if err != nil {
			return nil, err
		}
		if ok {
			res, err := ApplyJSONInterfaces(valuearray[1], data)
			if err != nil {
				return nil, err
			}
			return res, nil
		}

		if len(valuearray) > 2 {
			return opIf(valuearray[2:], data)
		}

		return nil, nil

	default: // A single value
		res, err := ApplyJSONInterfaces(value, data)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
