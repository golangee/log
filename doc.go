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

// Package log  is not another logger but a simple, clean and potentially
// dependency free logging facade. It combines ideas from Dave Cheney
// (https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
// and Rob Pike (https://github.com/golang/glog) to provide a simple yet efficient
// structured logging API.
//
// The default logger is created at package initialization time and
// if your application is executed from the IDE it uses the simple.PrintColored and otherwise
// simple.PrintStructured logger. Note, that the simple loggers use the standard library log.Print
// function and disables their time printing (log.SetFlags(0)).
package log
