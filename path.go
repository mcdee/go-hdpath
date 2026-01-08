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

// Path is a hierarchical derivation path with one element potentially set to "n".
type Path struct {
	elements []*element
}

// Parse parses a path.
func Parse(input string) (*Path, error) {
	elements := strings.Split(input, elementSeparator)
	if elements[0] != rootElement {
		return nil, ErrPathInvalid
	}
	elements = elements[1:]
	if len(elements) == 0 {
		return nil, ErrPathInvalid
	}

	res := &Path{
		elements: make([]*element, len(elements)),
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

// MustParse parses a path, panicking on an invalid input.
func MustParse(input string) *Path {
	res, err := Parse(input)
	if err != nil {
		panic(err)
	}

	return res
}

// Instance generates a fully resolved path.
func (p *Path) Instance(instance uint32) *Path {
	path := &Path{
		elements: make([]*element, len(p.elements)),
	}

	for i := range p.elements {
		path.elements[i] = p.elements[i].instance(instance)
	}

	return path
}

// Values provides the numeric values for the hierarchical path.
func (p *Path) Values() ([]uint32, error) {
	res := make([]uint32, len(p.elements))

	var err error
	for i := range p.elements {
		res[i], err = p.elements[i].value()
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

// String provides a string representation of the path.
func (p *Path) String() string {
	elements := make([]string, len(p.elements)+1)
	elements[0] = rootElement
	for i := range p.elements {
		elements[i+1] = p.elements[i].String()
	}

	return strings.Join(elements, elementSeparator)
}
