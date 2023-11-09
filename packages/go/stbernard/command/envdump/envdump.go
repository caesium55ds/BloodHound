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

package envdump

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	Name  = "envdump"
	Usage = "Dump your environment variables"
)

type cmd struct{}

func (s cmd) Name() string {
	return Name
}

func (s cmd) Usage() string {
	return Usage
}

func (s cmd) Run() error {
	for _, env := range os.Environ() {
		envTuple := strings.SplitN(env, "=", 2)
		fmt.Printf("%s: %s\n", envTuple[0], envTuple[1])
	}

	return nil
}

func Create() (cmd, error) {
	var (
		command    = cmd{}
		envdumpCmd = flag.NewFlagSet(Name, flag.ExitOnError)
	)

	envdumpCmd.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "%s\n\nUsage: %s %s [OPTIONS]\n", Usage, filepath.Base(os.Args[0]), Name)
	}

	if err := envdumpCmd.Parse(os.Args[2:]); err != nil {
		envdumpCmd.Usage()
		return command, fmt.Errorf("failed to parse %s command: %w", Name, err)
	} else {
		return command, nil
	}
}
