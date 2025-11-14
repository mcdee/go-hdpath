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

// Package hdpath managed hierarchical derivation path templates and instances.
//
// Templates are parsed from text, for example a path may be "m/44'/0'/n'/0/0" where:
// - m specifies the start of the path
// - a number specifies an index
// - a number followed by an apostrophe specifies a hardened index
// - n specifies an instance
//
// Templates are resolved to Paths by supplying an instance value, hence a Path is a fully-resolved Template.
// Path values can be obtained using the Values() function on a Path.
package hdpath

const (
	elementSeparator = "/"
	rootElement      = "m"
	variableElement  = "n"
	hardenedElement  = "'"
)
