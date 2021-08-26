package main

import (
	"os"

	"github.com/danangkonang/migration-go-cli/app/command"
	"github.com/danangkonang/migration-go-cli/app/helper"
)

func main() {
	arrCmd := os.Args[1:]
	if len(arrCmd) == 0 {
		helper.PrintHelper()
		return
	}
	runCmd()
}

func runCmd() {
	usrCmd := os.Args[1]
	switch usrCmd {
	/*
	* print helper cli
	 */
	case "-h":
		helper.PrintHelper()
	case "--help":
		helper.PrintHelper()

	/*
	* print version aplication cli
	 */
	case "-v":
		helper.PrintVersion()
	case "--version":
		helper.PrintVersion()

	/*
	* comand for init db migration and seeder directory
	* expanple: go run main.go init
	 */
	case "init":
		command.InitialV2()

	/*
	* comand for generate migration or seeder file
	* expanple: go run main.go create migration --name <table name>
	* file will be save to db/migration/<filename>
	* expanple: go run main.go create seeder --name <file name> --table <table name>
	* file will be save to db/seeder/<filename>
	 */
	case "create":
		command.MigrationSeederCreate()
	case "-c":
		command.MigrationSeederCreate()

	/*
	* comand for running migration or seeder
	* expanple: go run main.go run migration
	* expanple: go run main.go run seeder
	 */
	case "run":
		command.MigrationSeederRun()
	case "-r":
		command.MigrationSeederRun()

	/*
	* comand for delete all data in table
	* expanple: go run main.go reset
	 */
	case "reset":
		command.EmtySeederData()

	/*
	* comand for delete all tables in database
	* expanple: go run main.go drop
	 */
	case "drop":
		command.DropTableRun()
	default:
		helper.PrintHelper()
	}
}
