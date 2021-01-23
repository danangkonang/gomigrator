package command

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/execusion"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

func TypeMigration() {
	// dir := helper.MyRootDir()
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

func CreateMigration() {
	fullCmd := os.Args[1:]
	if len(fullCmd) < 3 {
		helper.PrintHelper()
		return
	}
	helper.CreateMigrationFile(helper.MyRootDir(), dirTableMigration)
	execusion.ReadeMiggrationFileInFolder(helper.MyRootDir())
}

func CreateSeeder() {
	fullCmd := os.Args[1:]
	if len(fullCmd) < 3 {
		helper.PrintHelper()
		return
	}
	helper.CreateSeedFile(helper.MyRootDir(), dirSeed)
	execusion.ReadeSeederFileInFolder(helper.MyRootDir())
}
