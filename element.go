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
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	hard = uint32(0x80000000)
)

type element struct {
	index    *uint32
	hardened bool
}

func (e *element) String() string {
	hardenedStr := ""
	if e.hardened {
		hardenedStr = hardenedElement
	}

	variableStr := variableElement
	if e.index != nil {
		variableStr = fmt.Sprintf("%d", *e.index)
	}

	return fmt.Sprintf("%s%s", variableStr, hardenedStr)
}

func parseElement(input string) (*element, error) {
	res := &element{
		index:    nil,
		hardened: false,
	}

	if input == "" {
		return nil, ErrElementMissing
	}

	if strings.HasSuffix(input, hardenedElement) {
		res.hardened = true
		input = input[:len(input)-1]
	}

	if input == variableElement {
		return res, nil
	}

	val, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return res, errors.Join(ErrElementInvalid, err)
	}
	value := uint32(val)

	res.index = &value

	return res, nil
}

func (e *element) instance(i uint32) *element {
	if e.index == nil {
		return &element{
			index:    &i,
			hardened: e.hardened,
		}
	}

	return e
}

func (e *element) value() (uint32, error) {
	if e.index == nil {
		return 0, ErrElementUnresolved
	}

	val := *e.index
	if e.hardened {
		val += hard
	}

	return val, nil
}
