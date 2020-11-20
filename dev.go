package log

import (
	"os"
	"strings"
)

// IsDevelopment evaluates the following sources:
//  - if any _INTELLIJ_* environment variable is defined, returns true
//  - XPC_SERVICE_NAME contains goland
//  - if APP_ENV or NODE_ENV environment variable is set to 'production' returns false, otherwise if specified at all
//    returns true
//  - if any VSCODE_* environment variable is defined, returns true
//  - otherwise returns false
func IsDevelopment() bool {
	for _, kv := range os.Environ() {
		if strings.HasPrefix(kv, "_INTELLIJ_") || strings.HasPrefix(kv, "VSCODE_") {
			return true
		}
	}

	if strings.Contains(os.Getenv("XPC_SERVICE_NAME"), "goland") {
		return true
	}

	nodeEnv := os.Getenv("APP_ENV")
	if strings.TrimSpace(nodeEnv) != "" {
		return nodeEnv != "production"
	}

	nodeEnv = os.Getenv("NODE_ENV")
	if strings.TrimSpace(nodeEnv) != "" {
		return nodeEnv != "production"
	}

	return false
}
