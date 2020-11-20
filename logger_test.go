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

//nolint: testpackage
package log_test

import (
	"context"
	"fmt"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"github.com/golangee/log/simple"
	"testing"
)

func TestNew(t *testing.T) {

	myLogger := log.NewLogger(ecs.Log("my.logger"))
	myLogger.Info(ecs.Msg("hello"))
	myLogger.Info(ecs.Msg("world"))

	fmt.Printf("=========\n")

	reqCtx := log.WithLogger(context.Background(), log.NewLogger(ecs.Log("my.request.logger")))
	log.FromContext(reqCtx).Info(ecs.Msg("from a request"))

	fmt.Printf("=========\n")

	logSomeStuff(log.LoggerFunc(ecs.WithName(ecs.WithTime(simple.Print), "my.logger")))
	logSomeStuff(log.LoggerFunc(ecs.WithName(ecs.WithTime(simple.PrintColored), "my.logger")))
	logSomeStuff(log.LoggerFunc(ecs.WithName(ecs.WithTime(simple.PrintStructured), "my.logger")))
}

func logSomeStuff(logger log.Logger) {
	logger.Info(ecs.Trace(), ecs.Msg("hello"))
	logger.Info(ecs.Debug(), ecs.Msg("hello"))
	logger.Info(ecs.Info(), ecs.Msg("hello"))
	logger.Info(ecs.Warn(), ecs.Msg("hello"))
	logger.Info(ecs.Error(), ecs.Msg("hello"), ecs.ErrStack())
	fmt.Print("\n\n---\n\n")
}
