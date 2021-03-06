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
	"fmt"
	"strings"

	"github.com/bannzai/constructor/generator"
	"github.com/bannzai/constructor/reader"
	"github.com/bannzai/constructor/structure"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	sourceFilePath      string
	destinationFilePath string
	templateFilePath    string
	structType          string
	ignoreFields        string
	packageName         string
}

var generateOptions = GenerateOptions{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate constructor functions",
	Long:  `constructor can be add constructor functions for each go struct. `,
	Run: func(command *cobra.Command, args []string) {
		generate()
	},
}

func generate() {
	ignoreFieldNames := []string{}
	if len(generateOptions.ignoreFields) > 0 {
		for _, splited := range strings.Split(generateOptions.ignoreFields, ",") {
			ignoreFieldNames = append(ignoreFieldNames, strings.Trim(splited, " "))
		}
	}

	generator.Constructor{
		TemplateReader:   reader.Template{},
		SourceCodeReader: reader.Code{},
		FileWriter:       generator.FileWriterImpl{},
	}.Generate(
		generateOptions.sourceFilePath,
		generateOptions.destinationFilePath,
		generateOptions.templateFilePath,
		generateOptions.structType,
		ignoreFieldNames,
		generateOptions.packageName,
	)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVar(&generateOptions.sourceFilePath, "source", "", "Source go file path")
	generateCmd.Flags().StringVar(&generateOptions.destinationFilePath, "destination", "", "Destination go file path")
	generateCmd.Flags().StringVar(&generateOptions.templateFilePath, "template", structure.TemplateFileName, fmt.Sprintf("Constructor functions format template file path. Default is ./%s", structure.TemplateFileName))
	generateCmd.Flags().StringVar(&generateOptions.structType, "type", "", "Specify struct about generated constructor function.")
	generateCmd.Flags().StringVar(&generateOptions.ignoreFields, "ignoreFields", "", "Not contains generated fields. It is list with commas. (e.g id,name,age")
	generateCmd.Flags().StringVar(&generateOptions.packageName, "package", "", "Package name for generated constructor.")
	requiredFlags := []string{"source", "destination", "package"}
	for _, name := range requiredFlags {
		if err := generateCmd.MarkFlagRequired(name); err != nil {
			panic(err)
		}
	}
}
