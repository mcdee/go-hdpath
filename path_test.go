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

package hdpath_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mcdee/go-hdpath"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		res   *hdpath.Path
		err   error
	}{
		{
			name:  "Empty",
			input: "",
			err:   hdpath.ErrPathInvalid,
		},
		{
			name:  "RootOnly",
			input: "m",
			err:   hdpath.ErrPathInvalid,
		},
		{
			name:  "SingleElement",
			input: "m/44",
			res:   hdpath.MustParse("m/44"),
		},
		{
			name:  "InvalidElement",
			input: "m/44'/0'/x/0/0",
			err:   hdpath.ErrElementInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := hdpath.Parse(test.input)
			if test.err != nil {
				if !errors.Is(err, test.err) {
					t.Fatalf("expected error %v, received error %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("received unexpected error %v", err)
				}
				if !reflect.DeepEqual(res, test.res) {
					t.Errorf("expected %v, received %v", test.res, res)
				}
			}
		})
	}
}

func TestString(t *testing.T) {
	zero := uint32(0)
	one := uint32(1)
	onetwothreefourfive := uint32(12345)
	tests := []struct {
		name     string
		p        *hdpath.Path
		instance *uint32
		res      string
	}{
		{
			name: "Unresolved",
			p:    hdpath.MustParse("m/44'/0'/n/0/0"),
			res:  "m/44'/0'/n/0/0",
		},
		{
			name: "UnresolvedHardened",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0"),
			res:  "m/44'/0'/n'/0/0",
		},
		{
			name:     "Zero",
			p:        hdpath.MustParse("m/44'/0'/n'/0/0"),
			instance: &zero,
			res:      "m/44'/0'/0'/0/0",
		},
		{
			name:     "One",
			p:        hdpath.MustParse("m/44'/0'/n'/0/0"),
			instance: &one,
			res:      "m/44'/0'/1'/0/0",
		},
		{
			name:     "OneTwoThreeFourFive",
			p:        hdpath.MustParse("m/44'/0'/n'/0/0"),
			instance: &onetwothreefourfive,
			res:      "m/44'/0'/12345'/0/0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := test.p
			if test.instance != nil {
				p = p.Instance(*test.instance)
			}

			res := p.String()
			if res != test.res {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}

func TestValues(t *testing.T) {
	onetwothreefourfive := uint32(12345)

	tests := []struct {
		name     string
		p        *hdpath.Path
		instance *uint32
		res      []uint32
		err      error
	}{
		{
			name: "Unresolved",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0"),
			err:  hdpath.ErrElementUnresolved,
		},
		{
			name:     "Resolved",
			p:        hdpath.MustParse("m/44'/0'/n'/0/0"),
			instance: &onetwothreefourfive,
			res: []uint32{
				44 + 0x80000000,
				0 + 0x80000000,
				12345 + 0x80000000,
				0,
				0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := test.p
			if test.instance != nil {
				p = p.Instance(*test.instance)
			}
			res, err := p.Values()
			if test.err != nil {
				if test.err != err {
					t.Fatalf("expected error %v, received error %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("received unexpected error %v", err)
				}
				if !reflect.DeepEqual(res, test.res) {
					t.Errorf("expected %v, received %v", test.res, res)
				}
			}
		})
	}
}
