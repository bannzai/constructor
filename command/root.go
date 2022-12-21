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
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "construtor",
	Short: "[constructor] can be generated constructor function for each struct",
	Long: `
[constructor] generate constructor functions for each struct.
When you execute "constructor generate [flags]", 
get constructor functions under the package.
You got "./constructor.tpl" after execute "constructor setup". 
This is default template for [constructor]. 
You can edit this template, If you customize generated files and pass it.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(command *cobra.Command, args []string) {
		command.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
