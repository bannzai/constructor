// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package command

import (
	"github.com/constructor/generator"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup will create ./constructor.yaml",
	Long: ` setup will create ./constructor.yaml.
	constructor.yaml is configuration file for "constructor generate".
	`,
	Run: func(command *cobra.Command, args []string) {
		generator.Template{}.Setup()
		generator.Yaml{}.Setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
