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
	instance *uint32
	hardened bool
}

func (e element) String(vals ...uint32) string {
	switch {
	case e.instance != nil:
		if e.hardened {
			return fmt.Sprintf("%d%s", *e.instance, hardenedElement)
		}

		return fmt.Sprintf("%d", *e.instance)
	case len(vals) == 0:
		if e.hardened {
			return variableElement + hardenedElement
		}

		return variableElement
	default:
		val := vals[0]
		if e.hardened {
			return fmt.Sprintf("%d%s", val, hardenedElement)
		}

		return fmt.Sprintf("%d", val)
	}
}

func parseElement(input string) (element, error) {
	res := element{
		instance: nil,
		hardened: false,
	}

	if input == "" {
		return res, ErrElementMissing
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

	res.instance = &value

	return res, nil
}

func (e element) value() (uint32, error) {
	switch {
	case e.instance == nil:
		return 0, ErrElementUnresolved
	case e.hardened:
		return *e.instance + hard, nil
	default:
		return *e.instance, nil
	}
}
