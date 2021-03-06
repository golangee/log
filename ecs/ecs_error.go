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

package ecs

import (
	"reflect"
	"runtime/debug"
	"strings"
)

// ErrStack is a factory to capture the current stack trace. This is quite expensive. The key is "error.stack_trace".
func ErrStack() Field {
	fullTrace := string(debug.Stack())
	lines := strings.Split(fullTrace, "\n")
	shortTrace := strings.Join(lines[6:], "\n")
	return Field{
		K: "error.stack_trace",
		V: shortTrace,
	}
}

// ErrMsg creates a field to note an error message. The key is "error.message". It captures err.String().
func ErrMsg(err error) Field {
	f := Field{
		K: "error.message",
	}

	if err != nil {
		f.V = err.Error()
	}

	return f
}

// ErrType creates a field to describe the go-name of the error type. The key is "error.type". When using
// fmt.Errorf this is mostly useless.
func ErrType(err error) Field {
	f := Field{
		K: "error.type",
	}

	if err != nil {
		f.V = reflect.TypeOf(err).String()
	}

	return f
}
