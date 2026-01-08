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
	"reflect"
	"testing"
)

func TestParseElement(t *testing.T) {
	ten := uint32(10)

	tests := []struct {
		name  string
		input string
		res   *element
		err   error
	}{
		{
			name: "Empty",
			res:  &element{},
			err:  ErrElementMissing,
		},
		{
			name:  "Value",
			input: "10",
			res: &element{
				index: &ten,
			},
		},
		{
			name:  "HardenedValue",
			input: "10'",
			res: &element{
				index:    &ten,
				hardened: true,
			},
		},
		{
			name:  "Variable",
			input: "n",
			res:   &element{},
		},
		{
			name:  "HardenedVariable",
			input: "n'",
			res: &element{
				hardened: true,
			},
		},
		{
			name:  "InvalidValue",
			input: "invalid",
			err:   ErrElementInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := parseElement(test.input)
			switch {
			case test.err != nil:
				if err == nil {
					t.Errorf("expected error %q, received no error", test.err)
				} else {
					if !errors.Is(err, test.err) {
						t.Errorf("expected error %q, received error %q", test.err, err.Error())
					}
				}
			default:
				if err != nil {
					t.Errorf("expected no error, received error %q", err.Error())
				} else {
					if !reflect.DeepEqual(res, test.res) {
						t.Errorf("expected %v, received %v", test.res, res)
					}
				}
			}
		})
	}
}

func TestElementString(t *testing.T) {
	ten := uint32(10)
	onetwothreefourfive := uint32(12345)

	tests := []struct {
		name     string
		e        *element
		instance *uint32
		res      string
	}{
		{
			name: "Value",
			e: &element{
				index: &ten,
			},
			res: "10",
		},
		{
			name: "HardenedValue",
			e: &element{
				index:    &ten,
				hardened: true,
			},
			res: "10'",
		},
		{
			name: "Variable",
			e:    &element{},
			res:  "n",
		},
		{
			name: "HardenedVariable",
			e: &element{
				hardened: true,
			},
			res: "n'",
		},
		{
			name:     "ResolvedVariable",
			e:        &element{},
			instance: &onetwothreefourfive,
			res:      "12345",
		},
		{
			name: "HardenedResolvedVariable",
			e: &element{
				hardened: true,
			},
			instance: &onetwothreefourfive,
			res:      "12345'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := test.e
			if test.instance != nil {
				e = test.e.instance(*test.instance)
			}

			res := e.String()
			if res != test.res {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}

func TestValue(t *testing.T) {
	ten := uint32(10)

	tests := []struct {
		name     string
		e        element
		instance uint32
		res      uint32
		err      error
	}{
		{
			name: "Ten",
			e: element{
				index: &ten,
			},
			res: 10,
		},
		{
			name: "TenHardened",
			e: element{
				index:    &ten,
				hardened: true,
			},
			res: 2147483658,
		},
		{
			name: "Unresolved",
			e: element{
				hardened: true,
			},
			err: ErrElementUnresolved,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := test.e.value()
			if test.err != nil {
				if test.err != err {
					t.Fatalf("expected error %v, received error %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("received unexpected error %v", err)
				}
				if res != test.res {
					t.Errorf("expected %v, received %v", test.res, res)
				}
			}
		})
	}
}
