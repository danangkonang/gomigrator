package command

import (
	"fmt"
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/execusion"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var dirMigration = "migration"
var dirDatabase = "migration/database"
var dirTableMigration = "migration/database/migration"
var dirSeed = "migration/database/seed"
var dirApp = "migration/app"
var dirCmd = "migration/app/command"
var dirExecusion = "migration/app/execusion"
var dirHelper = "migration/app/helper"
var dirConfig = "migration/app/config"
var dirTemplate = "migration/app/templates"

/*

	directory structure
	-------------------
	migration
		app
			helper
			config
			command
			execusion
			templates
		database
			migration
			seed
	--------------------
*/

func Initial() {

	makeDirectory(dirMigration)

	makeDirectory(dirDatabase)

	makeDirectory(dirTableMigration)

	makeDirectory(dirSeed)

	makeDirectory(dirApp)

	makeDirectory(dirCmd)

	makeDirectory(dirExecusion)

	makeDirectory(dirHelper)

	makeDirectory(dirConfig)
	makeDirectory(dirTemplate)

	thisDir := helper.MyRootDir()
	helper.CreateEnvFile()
	helper.CreateMainFile(thisDir, dirMigration)
	helper.CreateConfigFile(dirConfig)

	helper.CreateTemplateFile(thisDir, dirTemplate)

	// helper.CreateTypeCmd(thisDir, dirCmd)

	helper.CreatePrintTemplate(thisDir, dirHelper)
	// helper.CreateFileRootDirName(thisDir, dirHelper)

	// helper.CreateRunningMigrationFile(thisDir, dirExecusion)

	execusion.ReadeMiggrationFileInFolder(thisDir)
	execusion.ReadeSeederFileInFolder(thisDir)
}

func makeDirectory(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0700)
		fmt.Println("create " + name)
	}
}
