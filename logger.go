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
	"fmt"
	"log"
	"strings"
)

// Field is an alias to a key/value tuple, to break dependency.
type Field = struct {
	Key string
	Val interface{}
}

// Logger provides the different log levels and field definitions.
// It is kept as simple as possible and to avoid recursive dependency
// cycles by returns interfaces or concrete types.
type Logger interface {
	Trace(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

// Obj is a factory function for creating a Field.
func Obj(key string, val interface{}) Field {
	return Field{key, val}
}

//nolint: gochecknoglobals
var factory func(parent Logger, name string, fields ...Field) Logger = func(parent Logger, name string, fields ...Field) Logger {
	if sl, ok := parent.(simpleLogger); ok {
		if sl.name != "" {
			name = sl.name + "." + name
		}

		fields = append(sl.fields, fields...)
	}
	return simpleLogger{fields: fields, name: name}
}

// SetFactory just set a new root logger factory. The parent may be nil.
func SetFactory(f func(parent Logger, name string, fields ...Field) Logger) {
	factory = f
}

// New uses the factory to create a new root logger. Implementations are encouraged to
// return at least a reused default instance, if fields are empty.
func New(name string, fields ...Field) Logger {
	return With(nil, name, fields...)
}

// With is the same as New just with a parent logger. The factory may decide what to
// do, using a type assertion.
func With(parent Logger, name string, fields ...Field) Logger {
	return factory(parent, name, fields...)
}

// simpleLogger is a trivial implementation which just delegates to the go
// logger.
type simpleLogger struct {
	name   string
	fields []Field
}

func (d simpleLogger) With(fields ...Field) Logger {
	return simpleLogger{
		fields: append(d.fields, fields...),
	}
}

func (d simpleLogger) Trace(msg string, fields ...Field) {
	log.Println(cyan, "TRACE", reset, d.format(msg, fields...))
}

func (d simpleLogger) Debug(msg string, fields ...Field) {
	log.Println(magenta, "DEBUG", reset, d.format(msg, fields...))
}

func (d simpleLogger) Info(msg string, fields ...Field) {
	log.Println(blue, "INFO ", reset, d.format(msg, fields...))
}

func (d simpleLogger) Warn(msg string, fields ...Field) {
	log.Println(yellow, "WARN ", reset, d.format(msg, fields...))
}

func (d simpleLogger) Error(msg string, fields ...Field) {
	log.Println(red, "ERROR", reset, d.format(msg, fields...))
}

func (d simpleLogger) Panic(msg string, fields ...Field) {
	log.Panic(red, "PANIC", reset, d.format(msg, fields...))
}

func (d simpleLogger) Fatal(msg string, fields ...Field) {
	log.Fatal(red, "FATAL", reset, d.format(msg, fields...))
}

func (d simpleLogger) format(msg string, fields ...Field) string {
	sb := &strings.Builder{}
	if d.name != "" {
		sb.WriteString(gray)
		sb.WriteString("logger:")
		sb.WriteString(d.name)
		sb.WriteString(reset)
		sb.WriteString(" ")
	}
	sb.WriteString(msg)

	for _, f := range append(d.fields, fields...) {
		sb.WriteString(" ")
		sb.WriteString(gray)
		sb.WriteString(f.Key)
		sb.WriteString(":")
		sb.WriteString(reset)
		sb.WriteString(fmt.Sprintf("%v", f.Val))
	}

	return sb.String()
}
