// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"github.com/chrisbenson/epic/pkg/epic"
	//"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

// postgresqlCmd represents the postgresql command
var installPostgreSQLCmd = &cobra.Command{
	Use:   "postgresql",
	Short: "Installs a fresh Epic database on a PostgreSQL server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

epic install postgresql -u adminUser -p adminPassword -s pgServer

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verifyInstallFlags()
		fmt.Println("Creating Epic database.")
		appID, err := epic.InstallPostgreSQLDatabase(adminUser, adminPassword, server, database, epicPassword)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Epic database successfully created with this name: " + database)
			fmt.Println("Epic user created with password: " + epicPassword)
			fmt.Println("Epic application ID is: " + appID)
		}
		resetCreds()
	},
}

func init() {
	installCmd.AddCommand(installPostgreSQLCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postgresqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postgresqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	installPostgreSQLCmd.Flags().StringVarP(&adminUser, "adminUser", "u", "", "--adminUser admin, -u admin")
	installPostgreSQLCmd.Flags().StringVarP(&adminPassword, "adminPassword", "p", "", "--adminPassword password1, -p password1")
	installPostgreSQLCmd.Flags().StringVarP(&server, "server", "s", "", "--server server, -s server")
	installPostgreSQLCmd.Flags().StringVarP(&database, "database", "d", "", "--database database, -d database")
	installPostgreSQLCmd.Flags().StringVarP(&epicPassword, "epicPassword", "e", "", "--epicPassword password2, -e password2")
	installPostgreSQLCmd.MarkFlagRequired("adminUser")
	installPostgreSQLCmd.MarkFlagRequired("adminPassword")
	installPostgreSQLCmd.MarkFlagRequired("server")
	installPostgreSQLCmd.MarkFlagRequired("database")
	installPostgreSQLCmd.MarkFlagRequired("epicPassword")
}

func verifyInstallFlags() {

	var flagFail bool
	if adminUser == "" {
		flagFail = true
		fmt.Println("The required 'adminUser' flag has not been set. (e.g. --adminUser admin, -u admin)")
	}
	if adminPassword == "" {
		flagFail = true
		fmt.Println("The required 'adminPassword' flag has not been set. (e.g. --adminPassword password1, -p password1)")
	}
	if server == "" {
		flagFail = true
		fmt.Println("The required 'server' flag has not been set. (e.g. --server server, -s server)")
	}
	if database == "" {
		flagFail = true
		fmt.Println("The required 'database' flag has not been set. (e.g. --database database, -d database)")
	}
	if flagFail == true {
		fmt.Println("Correct usage is like this example: epic install postgresql --adminUser admin --adminPassword password1 --server server --database database")
		fmt.Println("Or this example: epic install postgresql -u admin -p password1 -s server -d database")
		os.Exit(1)
	}
}
