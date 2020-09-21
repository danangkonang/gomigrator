package cmd

import (
	"fmt"
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
		// CreateMigration()
		exec.RuningMigration()
		// fmt.Println("migration")
		break
	case "-m":
		// CreateMigration()
		// fmt.Println("-m")
		// exec.ReadeFileInFolder()
		exec.RuningMigration()
		break
	case "seeder":
		fmt.Println("seeder")
		// CreateSeeder()
		break
	case "-s":
		fmt.Println("-s")
		// CreateSeeder()
		break
	default:
		helper.PrintHelper()
	}
}

func TypeBack() {

}
