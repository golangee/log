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

// ServerAddress is ambiguous, may be the host name or the ip. The key is "server.address".
func ServerAddress(adr string) Field {
	return Field{
		Key: "server.address",
		Val: adr,
	}
}

// ServerAddress is the server domain. The key is "server.domain".
func ServerDomain(adr string) Field {
	return Field{
		Key: "server.domain",
		Val: adr,
	}
}

// ServerIp is the server ip. The key is "server.ip".
func ServerIp(adr string) Field {
	return Field{
		Key: "server.ip",
		Val: adr,
	}
}

// ServerPort is the server ip. The key is "server.port".
func ServerPort(port int) Field {
	return Field{
		Key: "server.port",
		Val: port,
	}
}
