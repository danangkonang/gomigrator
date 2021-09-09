package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/gomigrator/app/helper"
	"github.com/danangkonang/gomigrator/app/model"
)

func Init(app *model.Init) {
	if app.Driver == "" {
		app.Driver = "postgres"
	}
	if app.Host == "" {
		app.Host = "localhost"
	}
	if app.Port == 0 {
		app.Port = 5432
	}
	if app.Driver == "mysql" {
		app.Port = 3306
	}
	if app.DAtabase == "" {
		app.DAtabase = "default"
	}
	if app.User == "" {
		app.User = "user"
	}
	if app.Password == "" {
		app.Password = "password"
	}
	this_dir := helper.MyRootDir()
	// os.Exit(0)
	findTridparty(app)
	makeDirectoryV2(DirDb)
	makeDirectoryV2(DirMigration)
	makeDirectoryV2(DirSeeder)
	// helper.CreateBinFile(this_dir, DirDb)
	helper.CreateBinFileNew(this_dir, DirDb)
	helper.NewCreateZeroMigration(app)
	helper.CreateZeroSeeder()
	helper.NewCreateEnvFile(app)
}

func findTridparty(app *model.Init) {
	switch app.Driver {
	case "":
		pq := exec.Command("go", "get", "github.com/lib/pq")
		godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		pq.Run()
		godotenv.Run()
	case "postgres":
		pq := exec.Command("go", "get", "github.com/lib/pq")
		godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		pq.Run()
		godotenv.Run()
	case "mysql":
		pq := exec.Command("go", "get", "github.com/go-sql-driver/mysql")
		godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		pq.Run()
		godotenv.Run()
	default:
		fmt.Println("help me init")
		os.Exit(0)
	}
}
