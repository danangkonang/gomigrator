package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/migration-go-cli/app/helper"
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
	getTridparty()
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

	/*

	 ./.env

	*/
	helper.CreateEnvFile()

	/*

		./migration/migration.go

	*/
	helper.CreateMainFile(thisDir, dirMigration)

	/*

		./migration/app/config/config.go

	*/
	helper.CreateConfigFile(dirConfig)

	/*

		./migration/app/templates/template.go

	*/
	helper.CreateTemplateFile(thisDir, dirTemplate)

	// helper.CreateTypeCmd(thisDir, dirCmd)

	/*

		./migration/app/execution/type_execution.go

	*/
	helper.CreateTypeFileExecution()

	/*

		./migration/app/migration/type_migration.go

	*/
	helper.CreateTypeFileMigration()

	/*

		./migration/app/seed/type_seed.go

	*/
	helper.CreateTypeFileSeeder()

	/*

		./migration/app/helper/create_template.go

	*/
	helper.CreatePrintTemplate(thisDir, dirHelper)

	/*

		./migration/app/helper/get_time.go

	*/
	helper.CreateGetTime(thisDir, dirHelper)
	// helper.CreateFileRootDirName(thisDir, dirHelper)
	// helper.CreateRunningMigrationFile(thisDir, dirExecusion)

	/*

		./migration/app/execution/running_migration.go

	*/
	helper.ReadeMiggrationFileInFolder()

	/*

		./migration/app/execution/running_seeder.go

	*/
	helper.ReadeSeederFileInFolder()

	/*

		./migration/app/execution/down_table.go

	*/
	helper.ReadeDownFileInFolder()

	/*

		./migration/app/execution/drop_table.go

	*/
	helper.ReadeDropFileInFolder()
}

func makeDirectory(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0700)
		// pq := exec.Command("go", "get", "github.com/lib/pq")
		// mysql := exec.Command("go", "get", "github.com/go-sql-driver/mysql")
		// godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		// pq.Run()
		// mysql.Run()
		// godotenv.Run()
		fmt.Println("create " + name)
	}
}

func getTridparty() {
	pq := exec.Command("go", "get", "github.com/lib/pq")
	mysql := exec.Command("go", "get", "github.com/go-sql-driver/mysql")
	godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
	pq.Run()
	mysql.Run()
	godotenv.Run()
}
