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
	"context"
)

type ctxLogger struct{}

// WithLogger creates a new context with the given logger.
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

// FromContext returns the contained logger or a new root logger. Context may be nil.
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return NewLogger()
	}

	if logger, ok := ctx.Value(ctxLogger{}).(Logger); ok {
		return logger
	}

	return NewLogger()
}
