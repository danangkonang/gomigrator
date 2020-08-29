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
		cmd.Initial()
		break
	case "danang":
		cmd.CmdType(usrCmd)
		break
	case "create":
		cmd.TypeIs()
		break
	case "run":
		cmd.TypeIs()
		break
	case "back":
		cmd.TypeIs()
		break
	default:
		helper.PrintHelper()
	}
}
