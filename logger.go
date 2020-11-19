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
	Info(fields ...Field)
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
