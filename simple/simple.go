// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simple

import (
	"log"
)

// The Print logger just prints the fields in exactly the given order and converts the values to string using
// log.Print. You likely want to disable printing timestamps using log.SetFlags(0).
func Print(fields ...Field) {
	tmp := make([]interface{}, 0, len(fields))
	for i, field := range fields {
		tmp = append(tmp, field.Val)
		if i < len(fields)-1 {
			tmp = append(tmp, " ")
		}
	}

	log.Print(tmp...)
}
