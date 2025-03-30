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

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get configuration values.",
	Long: `Retrieve configuration value by specifyig a key.

Currently supproted keys:
  - token	: Retrieve the stored Notion API token.
	- db	: Retrieve the stored Notion database ID.`,
	Example: `  mnit get token
  mnit get db <user-defined-key>`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		switch key {
		case "token":
			token, err := service.GetToken()
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(token)
		case "db":
			if len(args) < 2 {
				return fmt.Errorf("missing user defined key.")
			}

			dbID, err := service.GetDBID(args[1])
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(dbID)
		default:
			return fmt.Errorf("invalid key: %s", key)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
