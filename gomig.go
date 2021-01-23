package main

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/command"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

// var dir = "database"

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
		command.TypeMigration() //go run main.go create migration [name file]
		break
	case "run":
		// command.TipeRun() //go run main.go run migration [name file]
		break
	case "back":
		command.TypeBack()
		break
		// case "run":
		// 	cmd.Run()
		// break
	default:
		helper.PrintHelper()
		// Execute()
	}
}
