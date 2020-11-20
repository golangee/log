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

const levelField = "log.level"

// Log applies the loggers name, like org.elasticsearch.bootstrap.Bootstrap.
func Log(name string) Field {
	return Field{
		Key: "log.logger",
		Val: name,
	}
}

// Trace creates the according level field. The key is "log.level".
func Trace() Field {
	return Field{
		Key: levelField,
		Val: "trace",
	}
}

// Debug creates the according level field. The key is "log.level".
func Debug() Field {
	return Field{
		Key: levelField,
		Val: "debug",
	}
}

// Info creates the according level field. The key is "log.level".
func Info() Field {
	return Field{
		Key: levelField,
		Val: "info",
	}
}

// Warn creates the according level field. The key is "log.level".
func Warn() Field {
	return Field{
		Key: levelField,
		Val: "warn",
	}
}

// Error creates the according level field. The key is "log.level".
func Error() Field {
	return Field{
		Key: levelField,
		Val: "error",
	}
}

// Fatal creates the according level field. The key is "log.level".
func Fatal() Field {
	return Field{
		Key: levelField,
		Val: "fatal",
	}
}

// Panic creates the according level field. The key is "log.level".
func Panic() Field {
	return Field{
		Key: levelField,
		Val: "panic",
	}
}
