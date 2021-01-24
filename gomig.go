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
		break
	case "--help":
		helper.PrintHelper()
		break
	case "-v":
		helper.PrintVersion()
		break
	case "--version":
		helper.PrintVersion()
		break
	case "init":
		command.Initial()
		break
	case "danang":
		command.TestingComand()
		break
	case "create":
		command.MigrationCreate() //go run main.go create migration [name file]
		break
	case "run":
		command.MigrationRun() //go run main.go run migration [name file]
		break
	case "back":
		command.MigrationUndo()
		break
	default:
		helper.PrintHelper()
	}
}