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
	"fmt"
	"strings"
)

// Path is a fully formed hierarchical path, with all values supplied.
type Path struct {
	elements []element
}

// Values provides the numeric values for the hierarchical path.
func (p Path) Values() []uint32 {
	res := make([]uint32, len(p.elements))

	var err error
	for i := range p.elements {
		res[i], err = p.elements[i].value()
		if err != nil {
			// This should never happen, because paths are only created by Template.Instance(), and that function ensures that
			// every element has a value assigned to it during creation.
			panic(fmt.Sprintf("malformed path element at %d", i))
		}
	}

	return res
}

// String provides a string representation of the path.
func (p Path) String() string {
	elements := make([]string, len(p.elements)+1)
	elements[0] = rootElement
	for i := range p.elements {
		elements[i+1] = p.elements[i].String()
	}

	return strings.Join(elements, elementSeparator)
}
