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
	"github.com/constructor/config"
	"github.com/constructor/input"
	"github.com/constructor/model"
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
	config.Configuration.YamlFilePath = YamlFilePathName
	yaml := input.YamlImpl{
		Argument: raw.Argument{
			YamlPath: configurationFilePath(),
		},
	}.Read()

	generateSources := []model.GenerateElementEachPackage{}
	for _, definition := range definitions(yaml) {
		templates := []raw.Template{}
		for _, path := range templateFilePaths(definition) {
			templates = append(templates, input.TemplateImpl{
				FilePath: path,
			}.Read())
		}

		codes := []raw.Code{}
		for _, path := range sourceCodeFilePaths(definition) {
			codes = append(codes, input.CodeImpl{
				FilePath: path,
			}.Read())
		}

		for _, template := range templates {
			generateSources = append(generateSources, model.GenerateElementEachPackage{
				Package:         definition.Package,
				Template:        template,
				Codes:           codes,
				DestinationPath: definition.DestinationPath,
			})
		}
	}

	for _, source := range generateSources {
		output.ConstructorImpl{source}.Generate()
	}
}

func definitions(yaml raw.Yaml) []raw.Definition {
	return []raw.Definition{}
}

func configurationFilePath() raw.Path {
	return "./" + YamlFilePathName
}

func templateFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
}

func sourceCodeFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
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
