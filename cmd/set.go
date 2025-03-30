/*
Copyright © 2025 team pascal

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/team-pascal/mnit/internal/service"
)

var setCmd = &cobra.Command{
	Use:   "set <key>",
	Short: "Set configuration values",
	Long: `Set configuration value by specifyig a key.

Currently supproted keys:
  - token	: Set the Notion API token. If a token has already stored, the token is updated.
	- db	: Set the Notion database ID. If the same key database id is stored, the database id is updated.`,
	Example: `  mnit set token <your-token>
  mnit set db <user-defined-key> <database id>`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		switch key {
		case "token":
			if len(args) < 2 {
				return fmt.Errorf("token is not specified")
			}

			token := args[1]

			if err := service.SetToken(token); err != nil {
				fmt.Println(err.Error())
			}
		case "db":
			if len(args) < 3 {
				return fmt.Errorf("missing user defined key or database id.")

			}

			dbKey := args[1]
			dbID := args[2]

			if err := service.SetDBID(dbKey, dbID); err != nil {
				fmt.Println(err.Error())
			}

		default:
			return fmt.Errorf("invalid key: %s", key)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
