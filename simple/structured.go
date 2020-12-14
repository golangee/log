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
	"github.com/golangee/log/field"
	"log"
	"runtime/debug"
)

// The PrintStructured logger takes the fields, removes duplicates (only the last is kept), and prints
// a json serialization as a single line using log.Print. The fields are sorted ascending by name. A special
// treatment is for message fields, which are simply fmt.Sprint'ed.
func PrintStructured(v ...interface{}) {
	fields := field.Fields(v...)
	tmp := make(map[string]interface{})
	for _, f := range fields {
		if f.K == "message" {
			s, ok := tmp[f.K]
			if ok {
				tmp[f.K] = fmt.Sprint(s, f.V)
			} else {
				tmp[f.K] = f.V
			}
		} else {
			tmp[f.K] = f.V
		}

	}

	buf, err := json.Marshal(tmp)
	if err != nil {
		log.Print("unable to marshal fields to json:", string(debug.Stack()), fmt.Sprint(fields))
		return
	}

	log.Print(string(buf))
}
