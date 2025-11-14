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
	"reflect"
	"testing"
)

func TestPathValues(t *testing.T) {
	ten := uint32(10)

	tests := []struct {
		name string
		path Path
		res  []uint32
	}{
		{
			name: "Empty",
			res:  []uint32{},
		},
		{
			name: "Good",
			path: Path{
				elements: []element{
					{
						instance: &ten,
					},
					{
						instance: &ten,
						hardened: true,
					},
				},
			},
			res: []uint32{10, 2147483658},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := (test.path.Values())
			if !reflect.DeepEqual(res, test.res) {
				t.Errorf("expected %v, received %v", test.res, res)
			}
		})
	}
}
