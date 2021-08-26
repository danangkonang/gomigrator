package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/migration-go-cli/app/helper"
)

var (
	DirDb        = "db"
	DirMigration = "db/migration"
	DirSeeder    = "db/seeder"
)

func InitialV2() {
	thisDir := helper.MyRootDir()
	getTridpartyV2()
	makeDirectoryV2(DirDb)
	makeDirectoryV2(DirMigration)
	makeDirectoryV2(DirSeeder)
	helper.CreateBinFile(thisDir, DirDb)
	helper.CreateZeroMigration()
	helper.CreateZeroSeeder()
	helper.CreateEnvFile()
}

func getTridpartyV2() {
	pq := exec.Command("go", "get", "github.com/lib/pq")
	godotenv := exec.Command("go", "get", "github.com/joho/godotenv")
	pq.Run()
	godotenv.Run()
}

func makeDirectoryV2(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		os.Mkdir(name, 0700)
		fmt.Println("create " + name)
	}
}
