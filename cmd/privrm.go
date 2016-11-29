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

	"github.com/akolb1/sentrytool/sentryapi"
	"github.com/spf13/cobra"
)

var privRevokeCmd = &cobra.Command{
	Use:     "revoke",
	Aliases: []string{"remove", "delete"},
	Short:   "revoke privilege",
	RunE:    revokePrivilege,
}

func revokePrivilege(cmd *cobra.Command, args []string) error {
	roleName, _ := cmd.Flags().GetString("role")
	if roleName == "" && len(args) == 0 {
		return errors.New("missing role name")
	}

	if roleName == "" {
		roleName = args[0]
	}

	action, _ := cmd.Flags().GetString("action")
	server, _ := cmd.Flags().GetString("server")
	database, _ := cmd.Flags().GetString("database")
	table, _ := cmd.Flags().GetString("table")
	column, _ := cmd.Flags().GetString("column")
	uri, _ := cmd.Flags().GetString("uri")
	scope, _ := cmd.Flags().GetString("scope")
	grant, _ := cmd.Flags().GetBool("grantoption")
	service, _ := cmd.Flags().GetString("service")

	priv := &sentryapi.Privilege{
		Action:      action,
		Server:      server,
		Database:    database,
		Table:       table,
		Column:      column,
		URI:         uri,
		Scope:       scope,
		GrantOption: grant,
		Service:     service,
	}

	client, err := getClient()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer client.Close()

	err = client.RevokePrivilege(roleName, priv)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func init() {
	privCmd.AddCommand(privRevokeCmd)
}