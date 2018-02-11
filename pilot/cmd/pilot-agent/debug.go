// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
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

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// TODO(nmittler): Remove this
	_ "github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	debugCmd = &cobra.Command{
		Use:   "local-config",
		Short: "Dump the local config",
		RunE: func(c *cobra.Command, args []string) error {
			var configFile os.FileInfo
			for configFile == nil {
				files, err := ioutil.ReadDir(configPath)
				if err != nil {
					return fmt.Errorf("unable to retrieve config file in %v: %v", configPath, err)
				}
				if len(files) == 1 {
					configFile = files[0]
				}
			}
			fileLocation := fmt.Sprintf("%v/%v", configPath, configFile.Name())
			bytes, err := ioutil.ReadFile(fileLocation)
			if err != nil {
				return fmt.Errorf("unable to read config file %q: %v", fileLocation, err)
			}
			fmt.Println(string(bytes))

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(debugCmd)
}
