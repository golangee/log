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
	"github.com/golangee/log/ecs"
	"github.com/golangee/log/simple"
	log2 "log"
)

// Field is an alias to a key/value tuple, to break dependency.
type Field = struct {
	Key string
	Val interface{}
}

// Logger provides the abstraction for logging.
// It is kept as simple as possible and to avoid recursive dependency
// cycles by using other interfaces or concrete types. It deliberately breaks with the conventional logger APIs, due
// to the following considerations:
//  * there are verbose developer specific logs which are not important or even bad for life system. You mostly
//    even want that there is no cost for the log parameter propagation, which would cause even more harm like
//    escaping values and heap-pressure, even if disabled. The only way to avoid this, is a guarded
//    compile time constant evaluation.
//  * anything else which is so important, that a developer is not sure to turn off in production, must not be guarded.
//    Instead it is up to the administrator or software engineer to filter through the log in a structured way. Any
//    error which does not kill your application, is just another kind of information.
//  * any error which does break your post-variants and you cannot continue, should be logged (again, still info)
//    either bail out with a runtime panic or an explicit os.Exit.
type Logger interface {
	// Info processes and prints the fields.
	Info(fields ...Field)
}

// LoggerFunc allows a function to become a Logger.
type LoggerFunc func(fields ...Field)

// Info prints the fields.
func (f LoggerFunc) Info(fields ...Field) {
	f(fields...)
}

//nolint: gochecknoglobals
var defaultFunc func(fields ...Field)

func init() {
	log2.SetFlags(0)
	if IsDevelopment() {
		defaultFunc = ecs.WithTime(simple.PrintColored)
	} else {
		defaultFunc = ecs.WithTime(simple.PrintStructured)
	}

}

// SetDefault just sets a delegate for NewLogger.
func SetDefault(f func(fields ...Field)) {
	defaultFunc = f
}

// NewLogger uses the factory to create a new logger. The given fields are prepended.
func NewLogger(fields ...Field) Logger {
	return LoggerFunc(func(f ...Field) {
		tmp := append(fields, f...)
		defaultFunc(tmp...)
	})
}

// WithFields just prepends the given fields before the actual logger field parameters will be passed.
func WithFields(logger Logger, fields ...Field) Logger {
	return LoggerFunc(func(f ...Field) {
		tmp := append(fields, f...)
		logger.Info(tmp...)
	})
}

// WithFunc returns a logger func which invokes the more field funcs and appends the given fields to it before
// invoking next.
func WithFunc(next func(fields ...Field), more ...func() Field) func(fields ...Field) {
	return func(fields ...Field) {
		tmp := make([]Field, 0, len(more))
		for _, f := range more {
			tmp = append(tmp, f())
		}

		tmp = append(tmp, fields...)
		next(tmp...)
	}
}

// V is just a shortcut for the field construction. V is a short version of Value.
func V(key string, val interface{}) Field {
	return Field{
		Key: key,
		Val: val,
	}
}
