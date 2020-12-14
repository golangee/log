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

package field

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Field is an interface to an explicit key/value tuple, to be clear about structured information.
type Field interface {
	// Key returns the unique key of this structured Field.
	Key() string
	// Value returns an arbitrary message, which itself may be structured.
	Value() interface{}
}

// DefaultField implements is a simple log.Field implementation.
type DefaultField struct {
	K string
	V interface{}
}

// Key returns the (unique) field name.
func (f DefaultField) Key() string {
	return f.K
}

// Value returns an arbitrary field value.
func (f DefaultField) Value() interface{} {
	return f.V
}

// String returns a debugging string serialization in the format <name>: <value> where value is serialized using
// the %v sprintf formatting directive.
func (f DefaultField) String() string {
	return fmt.Sprintf("%s: %v", f.K, f.V)
}

// Fields type casts the given interfaces or wraps them into multiple ECS compatible field types.
// It may return more fields than arguments, because it may logically parse or split an argument,
// like deriving error type and error message from an error.
func Fields(v ...interface{}) []DefaultField {
	res := make([]DefaultField, 0, len(v))
	for _, f := range v {
		switch t := f.(type) {
		case DefaultField:
			res = append(res, t)
		case *DefaultField:
			res = append(res, *t)
		case func() DefaultField:
			res = append(res, t())
		case Field:
			res = append(res, DefaultField{
				K: t.Key(),
				V: t.Value(),
			})
		case error:
			res = append(res, DefaultField{
				K: "error.message", // ecs standard
				V: t.Error(),
			})

			res = append(res, DefaultField{
				K: "error.type", // ecs standard
				V: reflect.TypeOf(t).String(),
			})
		case string:
			switch t {
			case "info":
				fallthrough
			case "trace":
				fallthrough
			case "debug":
				fallthrough
			case "fatal":
				fallthrough
			case "warn":
				res = append(res, DefaultField{
					K: "log.level", // ecs standard
					V: t,
				})
			default:
				if strings.HasPrefix(t, "http") {
					res = append(res, DefaultField{
						K: "url.path", // ecs standard
						V: t,
					})
				} else {
					res = append(res, DefaultField{
						K: "message", // ecs standard
						V: t,
					})
				}
			}
		case time.Time:
			res = append(res, DefaultField{
				K: "@timestamp", // ecs standard
				V: t.Format(time.RFC3339),
			})

		default:
			// everything else is also a message
			res = append(res, DefaultField{
				K: "message", // ecs standard
				V: fmt.Sprint(t),
			})
		}
	}

	return res
}
