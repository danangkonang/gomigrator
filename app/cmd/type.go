package cmd

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/exec"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

// cek is migration or seeder
func TypeMigration() {
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
	case "-m":
		CreateMigration()
		break
	case "seeder":
		CreateSeeder()
		break
	case "-s":
		CreateSeeder()
		break
	default:
		helper.PrintHelper()
	}
}

func TipeRun() {
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		exec.RuningMigration()
		break
	case "-m":
		exec.RuningMigration()
		break
	case "seeder":
		exec.RunSeeder()
		break
	case "-s":
		exec.RunSeeder()
		break
	default:
		helper.PrintHelper()
	}
}

func TypeBack() {

}
