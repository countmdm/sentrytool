// Copyright © 2016 Alex Kolbasov
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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// roleCmd represents the role command
var groupRemoveCmd = &cobra.Command{
	Use:     "revoke",
	Aliases: []string{"remove", "delete", "rm"},
	Short:   "Remove group from a role",
	Long: `Remove group from a role.
A role should be either specified with -role flag or be the first argument
followed by list of groups.

If role is specified with -role flag, arguments are group names to remove.`,
	Example: `
  group revoke admin_role admin_group finance_group
  group revoke -r admin_role admin_group finance_group`,
	RunE: removeGroupFromRole,
}

func removeGroupFromRole(cmd *cobra.Command, args []string) error {
	roleName, _ := cmd.Flags().GetString("role")
	if len(args) == 0 || (roleName == "" && len(args) == 1) {
		return errors.New("missing group name(s)")
	}

	groups := args
	if roleName == "" {
		roleName = args[0]
		groups = args[1:]
	}

	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer client.Close()

	// Verify that roleName is valid
	isValid, err := isValidRole(client, roleName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !isValid {
		return fmt.Errorf("role %s doesn't exist", roleName)
	}

	// Remove groups to the role
	if err = client.RemoveGroupsFromRole(roleName, groups); err != nil {
		fmt.Println(err)
		return nil
	}

	verbose := viper.Get(verboseOpt).(bool)
	if verbose {
		fmt.Println("removed groups from role", roleName)
	}

	return nil
}

func init() {
	groupCmd.AddCommand(groupRemoveCmd)
}
