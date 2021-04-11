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
	// case "back":
	// 	command.MigrationUndo()
	case "-b":
		command.MigrationUndo()
	case "down":
		command.Down("down") //delete rows
	case "drop":
		if len(os.Args[2:]) > 0 {
			// arguments := []string{}

			// for _, argumen := range os.Args[2:] {
			// 	arguments = append(arguments, argumen)
			// }
			// for i := 0; i < len(os.Args[2:]); i++ {
			// 	arguments = append(arguments, arguments[i+2])
			// }

			// fmt.Println(os.Args[2:])
			// arguments := strings.Join(os.Args[2:], " ")
			// fmt.Println(arguments)
			// command.Drop("drop table") //delete tables
		} else {
			command.Drop("drop") //delete tables
		}
	default:
		helper.PrintHelper()
	}
}
