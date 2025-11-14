// Copyright © 2025 Jim McDonald.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hdpath

import (
	"strings"
)

// Template is a hierarchical derivation path with one element potentially set to "n".
// Templates are resolved into paths using the Instance() function.
type Template struct {
	elements []element
}

// Parse parses a template.
func Parse(input string) (*Template, error) {
	elements := strings.Split(input, elementSeparator)
	if elements[0] != rootElement {
		return nil, ErrPathInvalid
	}
	elements = elements[1:]

	res := &Template{
		elements: make([]element, len(elements)),
	}

	var err error
	for i := range elements {
		res.elements[i], err = parseElement(elements[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// MustParse parses a template, panicking on an invalid input.
func MustParse(input string) *Template {
	res, err := Parse(input)
	if err != nil {
		panic(err)
	}

	return res
}

// Instance resolves a template into a path given an instance value.
func (t Template) Instance(val uint32) *Path {
	res := &Path{
		elements: make([]element, len(t.elements)),
	}

	for i := range t.elements {
		res.elements[i].hardened = t.elements[i].hardened
		res.elements[i].instance = t.elements[i].instance
		if res.elements[i].instance == nil {
			instance := val
			res.elements[i].instance = &instance
		}
	}

	return res
}

// String implements Stringer.
func (t Template) String() string {
	elements := make([]string, len(t.elements)+1)
	elements[0] = rootElement
	for i := range t.elements {
		elements[i+1] = t.elements[i].String()
	}

	return strings.Join(elements, elementSeparator)
}
