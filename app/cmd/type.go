package cmd

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

// cek is migration or seeder
func TypeIs() {
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		CreateMigration()
		break
	case "seeder":
		CreateSeeder()
		break
	default:
		helper.PrintHelper()
	}
}
