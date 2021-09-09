package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danangkonang/gomigrator/app/helper"
)

var (
	DirDb        = "db"
	DirMigration = "db/migration"
	DirSeeder    = "db/seeder"
)

func InitialV2() {
	// fmt.Println("panjang: ", len(os.Args))
	if len(os.Args) > 2 {
		if len(os.Args) == 4 {
			// fmt.Println(os.Args[3])
			thisDir := helper.MyRootDir()
			getTridpartyV2()
			makeDirectoryV2(DirDb)
			makeDirectoryV2(DirMigration)
			makeDirectoryV2(DirSeeder)
			helper.CreateBinFile(thisDir, DirDb)
			helper.CreateZeroMigration()
			helper.CreateZeroSeeder()
			helper.CreateEnvFile("mysql")
		} else {
			if os.Args[2] == "-d" || os.Args[2] == "--driver" {
				fmt.Println("unow driver")
			} else {
				err := fmt.Sprintf("unow %s", os.Args[2])
				fmt.Println(err)
			}
		}
	} else {
		// fmt.Println("driver: ", os.Args[1])
		thisDir := helper.MyRootDir()
		getTridpartyV2()
		makeDirectoryV2(DirDb)
		makeDirectoryV2(DirMigration)
		makeDirectoryV2(DirSeeder)
		helper.CreateBinFile(thisDir, DirDb)
		helper.CreateZeroMigration()
		helper.CreateZeroSeeder()
		helper.CreateEnvFile("postgres")
	}
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
