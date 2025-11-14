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

import "errors"

var (
	// ErrElementMissing is returned when a path element is empty.
	ErrElementMissing = errors.New("element missing")
	// ErrElementInvalid is returned when a path element is neither a value nor the instance value "n".
	ErrElementInvalid = errors.New("element invalid")
	// ErrElementUnresolved is returned when an attempt is made to obtain the value of an element that has not been fully defined.
	ErrElementUnresolved = errors.New("element unresolved")
	// ErrPathInvalid is returned when a string is not a valid hierarchical derivation path.
	ErrPathInvalid = errors.New("path invalid")
)
