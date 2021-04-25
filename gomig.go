package main

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/command"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
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
	case "-h":
		helper.PrintHelper()
	case "--help":
		helper.PrintHelper()
	case "-v":
		helper.PrintVersion()
	case "--version":
		helper.PrintVersion()
	case "init":
		command.Initial()
	case "create":
		command.MigrationCreate() //go run main.go create migration [name file]
	case "-c":
		command.MigrationCreate() //go run main.go create migration [name file]
	case "run":
		// command.MigrationRun() //go run main.go run migration [name file]
		command.MigrationsRun()
	case "-r":
		command.MigrationsRun() //go run main.go run migration [name file]
	case "down":
		command.EmtySeederData() // delete data seeder
	case "drop":
		command.DropTableRun() // delete table
	default:
		helper.PrintHelper()
	}
}
