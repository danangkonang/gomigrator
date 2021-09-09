package command

import (
	"fmt"
	"os"

	"github.com/danangkonang/gomigrator/app/helper"
)

func MigrationSeederCreate() {
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		_, err := os.Stat("db/migration")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		CreateMigrationTableV2()
	case "seeder":
		_, err := os.Stat("db/seeder")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		CreateSeederTableV2()
	case "-h":
		helper.MultyPrintHelper("create")
	case "--help":
		helper.MultyPrintHelper("create")
	default:
		helper.PrintHelper()
	}
}
