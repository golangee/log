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
	"time"
)

// Msg creates a message field. The key is "message". It contains an optimized view for humans.
func Msg(msg string) Field {
	return Field{
		Key: "message",
		Val: msg,
	}
}

// Time is arguable, at the one hand the systems logger already captures the time, however when sending
// and processing structured logs it may be useful to include an applications timestamp. The key is "@timestamp"
// and the format is RFC3339. It describes the Date/time when the event originated.
func Time() Field {
	return Field{
		Key: "@timestamp",
		Val: time.Now().Format(time.RFC3339),
	}
}

// Tags is a list of keywords used to tag each event. The key is "tags".
func Tags(tags ...string) Field {
	return Field{
		Key: "tags",
		Val: tags,
	}
}

// Labels are un-nestable custom key/value pairs. The key is "labels".
func Labels(labels map[string]string) Field {
	return Field{
		Key: "tags",
		Val: labels,
	}
}
