// Copyright © 2016 Jessica Ete <kohrVid@zoho.com>
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
	textGoesFunny "github.com/kohrVid/textGoesFunny/app"
	"github.com/spf13/cobra"
)

// _1337Cmd represents the _1337 command
var _1337Cmd = &cobra.Command{
	Use:   "_1337",
	Short: "Display 1337 characters",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		HelpCommand(cmd, args)
		_1337 := textGoesFunny.L337
		textGoesFunny.HandleString(args[0], _1337)
	},
}

func init() {
	RootCmd.AddCommand(_1337Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// _1337Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// _1337Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
