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

package cmd

import (
	"context"

	"github.com/constructor/config"
	"github.com/constructor/input"
	"github.com/constructor/output"
	"github.com/constructor/raw"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate constructor functions",
	Long: `When use "generate" command, some constructor functions are generated.
constructor generate [/path/to/package] [-c(--config) constructor.yaml].
`,
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	ctx := context.Background()
	config.Configuration.YamlFilePath = YamlFilePathName
	output.ConstructorImpl{
		YamlReader: input.YamlImpl{
			Argument: raw.Argument{
				YamlPath: YamlFilePathName,
			},
		},
		TemplateReader:   input.TemplateImpl{},
		SourceCodeReader: input.CodeImpl{},
	}.Generate(ctx)
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
