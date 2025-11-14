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
	"reflect"
	"testing"

	"github.com/mcdee/go-hdpath"
)

func TestPathString(t *testing.T) {
	tests := []struct {
		name string
		p    *hdpath.Path
		vals []uint32
		res  string
	}{
		{
			name: "Root",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0").Instance(0),
			res:  "m/44'/0'/0'/0/0",
		},
		{
			name: "Variable",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0").Instance(1),
			res:  "m/44'/0'/1'/0/0",
		},
		{
			name: "ResolvedVariable",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0").Instance(12345),
			res:  "m/44'/0'/12345'/0/0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.p.String()
			if res != test.res {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		name string
		p    *hdpath.Path
		res  []uint32
	}{
		{
			name: "Static",
			p:    hdpath.MustParse("m/44'/0'/n'/0/0").Instance(12345),
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
			res := test.p.Values()
			if !reflect.DeepEqual(res, test.res) {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}
