package main

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/cmd"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var dir = "database"

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
		cmd.Initial() //create .env config database
		break
	case "danang":
		cmd.TypeMigration()
		break
	case "create":
		cmd.TypeMigration() //go run main.go create migration [name file]
		break
	case "run":
		cmd.TipeRun() //go run main.go run migration [name file]
		break
	case "back":
		cmd.TypeBack()
		break
		// case "run":
		// 	cmd.Run()
		// break
	default:
		helper.PrintHelper()
		// Execute()
	}
}
