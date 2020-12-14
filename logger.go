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
	"github.com/golangee/log/field"
	"github.com/golangee/log/simple"
	log2 "log"
)

// Field is an interface to an explicit key/value tuple, to be clear about structured information.
type Field interface {
	// Key returns the unique key of this structured Field.
	Key() string
	// Value returns an arbitrary message, which itself may be structured.
	Value() interface{}
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
	// Println processes and prints the arguments as fields. The interpretation and formatting depends on the
	// concrete implementation and may range from fmt.Println over log.Println to a full structured logger.
	// Implementations are encouraged to type-switch on each field.
	Println(fields ...interface{})
}

// LoggerFunc allows a function to become a Logger.
type LoggerFunc func(fields ...interface{})

// Println prints the fields by delegating to the function itself.
func (f LoggerFunc) Println(fields ...interface{}) {
	f(fields...)
}

//nolint: gochecknoglobals
var defaultFunc LoggerFunc

func init() {
	log2.SetFlags(0)
	if IsDevelopment() {
		defaultFunc = ecs.WithTime(simple.PrintColored)
	} else {
		defaultFunc = ecs.WithTime(simple.PrintStructured)
	}

}

// SetDefault just sets a delegate for NewLogger.
func SetDefault(f func(fields ...interface{})) {
	defaultFunc = f
}

// NewLogger uses the factory to create a new logger. The given fields are prepended.
func NewLogger(fields ...interface{}) Logger {
	if len(fields) == 0 {
		return defaultFunc
	}

	return LoggerFunc(func(f ...interface{}) {
		tmp := append(fields, f...)
		defaultFunc(tmp...)
	})
}

// Println delegates directly to the default configured logger. See also Logger, NewLogger and SetDefault.
func Println(fields ...interface{}) {
	defaultFunc(fields...)
}

// WithFields just prepends the given fields before the actual logger field parameters will be passed.
func WithFields(logger Logger, fields ...interface{}) Logger {
	if len(fields) == 0 {
		return logger
	}

	return LoggerFunc(func(f ...interface{}) {
		tmp := append(fields, f...)
		logger.Println(tmp...)
	})
}

// WithFunc returns a logger func which invokes the more field funcs and appends the given fields to it before
// invoking next.
func WithFunc(next func(fields ...interface{}), more ...func() interface{}) func(fields ...interface{}) {
	return func(fields ...interface{}) {
		tmp := make([]interface{}, 0, len(more))
		for _, f := range more {
			tmp = append(tmp, f())
		}

		tmp = append(tmp, fields...)
		next(tmp...)
	}
}

// V is just a shortcut for a field construction. V is a short version of Value.
func V(key string, val interface{}) field.DefaultField {
	return field.DefaultField{
		K: key,
		V: val,
	}
}
