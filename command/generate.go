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
	"github.com/bannzai/constructor/generator"
	"github.com/bannzai/constructor/reader"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	sourceFilePath      string
	destinationFilePath string
	ignoreFields        string
	templateFilePath    string
}

var generateOptions = GenerateOptions{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate constructor functions",
	Long: `When use "generate" command, some constructor functions are generated.
constructor generate [/path/to/package] [-c(--config) constructor.yaml].
`,
	Run: func(command *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	generator.Constructor{
		TemplateReader:   reader.Template{},
		SourceCodeReader: reader.Code{},
		FileWriter:       generator.FileWriter{},
		FilePathFetcher:  generator.Glob{},
	}.Generate()
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateOptions.sourceFilePath, "source", "", "", "Source go file path")
	generateCmd.Flags().StringVarP(&generateOptions.destinationFilePath, "destination", "", "", "Destination go file path")
	generateCmd.Flags().StringVarP(&generateOptions.ignoreFields, "ignoreFields", "", "", "Not contains generated fields. It is list with commas. (e.g id,name,age")
	generateCmd.Flags().StringVarP(&generateOptions.templateFilePath, "tempalte", "", "", "Constructor functions format template file path")
}
