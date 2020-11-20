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
	"encoding/json"
	"fmt"
	"log"
	"runtime/debug"
)

// The PrintStructured logger takes the fields, removes duplicates (only the last is kept), and prints
// a json serialization as a single line using log.Print. The fields are sorted ascending by name.
func PrintStructured(fields ...Field) {
	tmp := make(map[string]interface{})
	for _, field := range fields {
		tmp[field.Key] = field.Val
	}

	buf, err := json.Marshal(tmp)
	if err != nil {
		log.Print("unable to marshal fields to json:", string(debug.Stack()), fmt.Sprint(fields))
		return
	}

	log.Print(string(buf))
}
