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
	case "danang":
		command.TestingComand()
	case "create":
		command.MigrationCreate() //go run main.go create migration [name file]
	case "-c":
		command.MigrationCreate() //go run main.go create migration [name file]
	case "run":
		command.MigrationRun() //go run main.go run migration [name file]
	case "-r":
		command.MigrationRun() //go run main.go run migration [name file]
	case "-b":
		command.MigrationUndo()
	case "down":
		command.Down("down", os.Args[2:]) //delete tables
	case "drop":
		command.Drop("drop", os.Args[2:]) //delete tables
	default:
		helper.PrintHelper()
	}
}
