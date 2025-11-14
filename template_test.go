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
	"testing"

	"github.com/mcdee/go-hdpath"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		err   error
	}{
		{
			name:  "RootMissing",
			input: "44'/0'/n'/0/0",
			err:   hdpath.ErrPathInvalid,
		},
		{
			name:  "RootInvalid",
			input: "x/44'/0'/n'/0/0",
			err:   hdpath.ErrPathInvalid,
		},
		{
			name:  "ElementMissing",
			input: "m/44'/0'//0/0",
			err:   hdpath.ErrElementMissing,
		},
		{
			name:  "ElementBad",
			input: "m/44'/0'/x/0/0",
			err:   hdpath.ErrElementInvalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := hdpath.Parse(test.input)
			switch {
			case test.err != nil:
				switch err {
				case nil:
					t.Errorf("expected error %v, received no error", test.err)
				default:
					if !errors.Is(err, test.err) {
						t.Errorf("expected error %v, received error %v", test.err, err)
					}
				}
			default:
				switch {
				case err != nil:
					t.Errorf("expected no error, received error %v", err)
				default:
					if res.String() != test.input {
						t.Errorf("expected %v, received %v", test.input, res.String())
					}
				}
			}
		})
	}
}

func TestTemplateString(t *testing.T) {
	tests := []struct {
		name string
		t    *hdpath.Template
		res  string
	}{
		{
			name: "Root",
			t:    hdpath.MustParse("m/44'/0'/n'/0/0"),
			res:  "m/44'/0'/n'/0/0",
		},
		{
			name: "Variable",
			t:    hdpath.MustParse("m/44'/0'/1'/0/0"),
			res:  "m/44'/0'/1'/0/0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := test.t.String()
			if res != test.res {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}
