// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/specterops/bloodhound/packages/go/stbernard/command"
	"github.com/specterops/bloodhound/packages/go/stbernard/environment"
	"golang.org/x/exp/slog"
)

func main() {
	setLogger()

	if cmd, err := command.ParseCLI(); err != nil {
		slog.Error("Failed to parse CLI", "error", err)
		os.Exit(1)
	} else if err := cmd.Run(); err != nil {
		slog.Error("Failed to run command", "command", cmd.Name(), "error", err)
		os.Exit(1)
	} else {
		slog.Info("Command completed successfully", "command", cmd.Name())
	}
}

func setLogger() {
	var logLevel slog.Level

	envLvl := os.Getenv(environment.LogLevel.Env())

	switch envLvl {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))
}
