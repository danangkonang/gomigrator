package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/migrasion-go-cli/app/execusion"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var dirMigration = "migration"
var dirDatabase = "migration/database"
var dirTableMigration = "migration/database/migration"
var dirSeed = "migration/database/seed"
var dirApp = "migration/app"

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
	helper.CreateGetTime(thisDir, dirHelper)
	// helper.CreateFileRootDirName(thisDir, dirHelper)
	// helper.CreateRunningMigrationFile(thisDir, dirExecusion)

	execusion.ReadeMiggrationFileInFolder()
	execusion.ReadeSeederFileInFolder()

	helper.ReadeDownFileInFolder()
	helper.ReadeDropFileInFolder()
}

func makeDirectory(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0700)
		pq := exec.Command("go", "get", "github.com/lib/pq")
		mysql := exec.Command("go", "get", "github.com/go-sql-driver/mysql")
		godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		pq.Run()
		mysql.Run()
		godotenv.Run()
		fmt.Println("create " + name)
	}
}
