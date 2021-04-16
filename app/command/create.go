package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/execusion"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

func MigrationCreate() {
	//  CREATE
	fullCommand := os.Args[1:]
	if len(fullCommand) == 1 {
		helper.PrintHelper()
		return
	}
	subCommand := os.Args[2]
	switch subCommand {
	case "migration":
		_, err := os.Stat("migration/database/migration")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		CreateMigration()
	case "seeder":
		_, err := os.Stat("migration/database/seed")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		CreateSeeder()
	case "-h":
		helper.PrintHelperCreate()
	case "--help":
		helper.PrintHelperCreate()
	default:
		helper.PrintHelper()
	}
}

func CreateMigration() {
	var migration helper.Migration
	if len(os.Args[3:]) == 0 {
		fmt.Println("migration tidak ada argumen ")
		return
	}
	if len(os.Args[3:]) > 2 {
		fmt.Println(string("\033[31m"), "argumen "+os.Args[5]+" tidak dikenal")
		os.Exit(0)
	}
	if os.Args[3] == "-t" || os.Args[3] == "table" {
		fmt.Println(len(os.Args[3:]))
		migration.DirMigration = dirTableMigration
		migration.TableMigration = os.Args[4]
		helper.CreateMigrationFile(&migration)
		execusion.ReadeMiggrationFileInFolder()
		helper.ReadeDownFileInFolder()
		helper.ReadeDropFileInFolder()
	} else {
		fmt.Println(string("\033[31m"), "argumen "+os.Args[3]+" tidak dikenal")
		os.Exit(0)
	}
}

func CreateSeeder() {
	var seed helper.Seeder
	if len(os.Args[3:]) == 0 {
		fmt.Println("tidak ada argumen")
		return
	}

	for i := 0; i < len(os.Args[3:]); i++ {
		if os.Args[i+3] == "-n" {
			if len(os.Args[i+3:]) < 2 {
				fmt.Println("nama file tidak boleh kosong")
				return
			}
			seed.Filename = os.Args[i+4] // set nama file
		}
		if os.Args[i+3] == "-t" {
			if len(os.Args[i+3:]) < 2 {
				fmt.Println("nama tabel tidak boleh kosong")
				return
			}
			seed.TableName = os.Args[i+4] // set nama file
		}
	}
	seed.DirSeeder = dirSeed
	if seed.Filename == "" {
		fmt.Println("nama file tidak boleh kosong")
		return
	}
	if seed.TableName == "" {
		fmt.Println("nama tabel tidak boleh kosong")
		return
	}

	var file, _ = os.OpenFile(helper.MyRootDir(), os.O_RDWR, 0644)
	defer file.Close()
	// fmt.Println(file)

	files, err := ioutil.ReadDir("migration/database/seed")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		list := strings.Split(filename, "_")
		name := list[0]
		if name == seed.Filename {
			fmt.Println("nama file harus unik")
			os.Exit(0)
		}
	}

	helper.CreateSeedFile(&seed)
	execusion.ReadeSeederFileInFolder()
	helper.ReadeDownFileInFolder()
	helper.ReadeDropFileInFolder()
}
