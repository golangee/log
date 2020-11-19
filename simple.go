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

package log

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// simpleLogger is a trivial implementation which just delegates directly to stdout.
type simpleLogger struct {
	name   string
	fields []Field
}

func (d simpleLogger) With(fields ...Field) Logger {
	return simpleLogger{
		fields: append(d.fields, fields...),
	}
}

func (d simpleLogger) Info(fields ...Field) {
	fmt.Println(format(true, fields...))
}

func format(colorize bool, fields ...Field) string {
	sb := &strings.Builder{}
	sb.WriteByte('{')
	for i, field := range fields {
		needsReset := false
		if colorize {
			switch field.Key {
			case levelField:
				if str, ok := field.Val.(string); ok {
					needsReset = true
					switch str {
					case "debug":
						sb.WriteString(magenta)
					case "info":
						sb.WriteString(blue)
					case "trace":
						sb.WriteString(cyan)
					case "warn":
						sb.WriteString(yellow)
					default:
						sb.WriteString(red)
					}
				}

			}
		}

		sb.WriteString(strconv.Quote(field.Key))
		sb.WriteString(":")

		buf, err := json.Marshal(field.Val)
		if err != nil {
			sb.WriteString(strconv.Quote(err.Error()))
		} else {
			sb.Write(buf)
		}

		if needsReset {
			sb.WriteString(reset)
		}

		if i < len(fields)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteByte('}')
	return sb.String()
}
