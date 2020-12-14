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
	"fmt"
	"github.com/golangee/log/field"
	"log"
	"strings"
)

const intendTrace = len("2020-11-20T10:54:11+01:00")

// The PrintColored logger just prints the fields in exactly the given order and converts the values to string using
// log.Print. You likely want to disable printing timestamps using log.SetFlags(0).
// The console print is scattered with color commands and probably only nice for your developer
// machine.
func PrintColored(v ...interface{}) {
	tmp := make([]interface{}, 0, len(v))
	fields := field.Fields(v...)
	messageColor := ""
	for i, field := range fields {
		needsReset := false

		switch field.K {
		case "log.level":
			if str, ok := field.V.(string); ok {
				needsReset = true
				switch str {
				case "debug":
					messageColor = magenta
				case "info":
					messageColor = blue
				case "trace":
					messageColor = cyan
				case "warn":
					messageColor = yellow
				default:
					messageColor = red
				}
				if str, ok := field.V.(string); ok {
					field.V = strings.ToUpper(str)
				}

				tmp = append(tmp, messageColor)
			}
		case "@timestamp":
			needsReset = true
			tmp = append(tmp, cyan)
		case "error.stack_trace":
			indent := &strings.Builder{}
			for i := 0; i < intendTrace; i++ {
				indent.WriteByte(' ')
			}

			needsReset = true
			tmp = append(tmp, red)
			if str, ok := field.V.(string); ok {
				field.V = strings.ReplaceAll(str, "\n", "\n"+indent.String()+red)
			}
		case "message":
			if messageColor != "" {
				needsReset = true
				tmp = append(tmp, messageColor)
			}
		}

		tmp = append(tmp, fmt.Sprint(field.V))

		if needsReset {
			tmp = append(tmp, reset)
		}

		if i < len(fields)-1 {
			tmp = append(tmp, " ")
		}
	}

	log.Print(tmp...)
}
