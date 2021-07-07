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
var dirApp = "migration/core"

var dirExecusion = "migration/core/execusion"
var dirHelper = "migration/core/helper"
var dirConfig = "migration/core/config"
var dirTemplate = "migration/core/templates"

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

	// CopyGomig()

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

		./migration/core/config/config.go

	*/
	helper.CreateConfigFile(dirConfig)

	/*

		./migration/core/templates/template.go

	*/
	helper.CreateTemplateFile(thisDir, dirTemplate)

	// helper.CreateTypeCmd(thisDir, dirCmd)

	/*

		./migration/core/execution/type_execution.go

	*/
	helper.CreateTypeFileExecution()

	/*

		./migration/core/migration/type_migration.go

	*/
	helper.CreateTypeFileMigration()

	/*

		./migration/core/seed/type_seed.go

	*/
	helper.CreateTypeFileSeeder()

	/*

		./migration/core/helper/create_template.go

	*/
	helper.CreatePrintTemplate(thisDir, dirHelper)

	/*

		./migration/core/helper/get_time.go

	*/
	helper.CreateGetTime(thisDir, dirHelper)
	// helper.CreateFileRootDirName(thisDir, dirHelper)
	// helper.CreateRunningMigrationFile(thisDir, dirExecusion)

	/*

		./migration/core/execution/running_migration.go

	*/
	helper.ReadeMiggrationFileInFolder()

	/*

		./migration/core/execution/running_seeder.go

	*/
	helper.ReadeSeederFileInFolder()

	/*

		./migration/core/execution/down_table.go

	*/
	helper.ReadeDownFileInFolder()

	/*

		./migration/core/execution/drop_table.go

	*/
	helper.ReadeDropFileInFolder()
}

// func CopyGomig() {
// 	gomig := exec.Command("cp", "gomig", "migration")
// 	gomig.Run()
// }

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
	// mysql := exec.Command("go", "get", "github.com/go-sql-driver/mysql")
	godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
	pq.Run()
	// mysql.Run()
	godotenv.Run()
}
