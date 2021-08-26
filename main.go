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
	case "-h":
		helper.PrintHelper()
	case "--help":
		helper.PrintHelper()
	case "-v":
		helper.PrintVersion()
	case "--version":
		helper.PrintVersion()
	case "init":
		command.InitialV2()
	case "create":
		command.MigrationSeederCreate() //go run main.go create migration [name file]
	case "-c":
		command.MigrationSeederCreate() //go run main.go create migration [name file]
	case "run":
		command.MigrationSeederRun()
	case "-r":
		command.MigrationSeederRun() //go run main.go run migration [name file]
	case "reset":
		command.EmtySeederData() // delete data seeder
	case "drop":
		command.DropTableRun() // delete table
	default:
		helper.PrintHelper()
	}
}
