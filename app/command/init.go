package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/migration-go-cli/app/helper"
	"github.com/danangkonang/migration-go-cli/app/model"
)

func Init(app *model.Init) {
	if app.Driver == "" {
		app.Driver = "postgresql"
	}
	if app.Host == "" {
		app.Host = "localhost"
	}
	if app.Port == 0 {
		app.Port = 5432
	}
	if app.DAtabase == "" {
		app.DAtabase = "dafault"
	}
	if app.User == "" {
		app.User = "root"
	}
	if app.Password == "" {
		app.Password = "root"
	}
	this_dir := helper.MyRootDir()
	findTridparty(app)
	makeDirectoryV2(DirDb)
	makeDirectoryV2(DirMigration)
	makeDirectoryV2(DirSeeder)
	helper.CreateBinFile(this_dir, DirDb)
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
	case "postgresql":
		pq := exec.Command("go", "get", "github.com/lib/pq")
		godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
		pq.Run()
		godotenv.Run()
	case "mysql":
	default:
		fmt.Println("help me")
		os.Exit(0)
	}
}
