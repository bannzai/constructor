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

	"github.com/constructor/generator"
	"github.com/constructor/reader"
	"github.com/constructor/structure"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	yamlFilePath string
}

var generateOptions = GenerateOptions{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate constructor functions",
	Long: `When use "generate" command, some constructor functions are generated.
constructor generate [/path/to/package] [-c(--config) constructor.yaml].
`,
	Run: func(cmd *cobra.Command, args []string) {
		generate(generateOptions.yamlFilePath)
	},
}

func generate(yamlFilePath string) {
	ctx := context.Background()
	generator.Constructor{
		YamlReader: reader.YamlImpl{
			Argument: structure.Argument{
				YamlPath: yamlFilePath,
			},
		},
		TemplateReader:   reader.TemplateImpl{},
		SourceCodeReader: reader.CodeImpl{},
	}.Generate(ctx)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateOptions.yamlFilePath, "configure", "c", structure.YamlFilePathName, "Specify configure file")
}
