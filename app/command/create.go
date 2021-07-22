package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/helper"
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
		_, err := os.Stat("db/migration")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		// CreateMigration()
		CreateMigrationTableV2()
	case "seeder":
		_, err := os.Stat("db/seeder")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		CreateSeeder()
	case "-h":
		helper.MultyPrintHelper("create")
	case "--help":
		helper.MultyPrintHelper("create")
	default:
		helper.PrintHelper()
	}
}

func CreateMigration() {
	var migration helper.Migration
	if len(os.Args[3:]) == 0 {
		// fmt.Println("migration tidak ada argumen ")
		fmt.Println(string("\033[31m"), "table name not found")
		fmt.Println(string("\033[00m"), "please run: gomig create --help")
		return
	}
	if len(os.Args[3:]) > 2 {
		fmt.Println(string("\033[31m"), "argumen "+os.Args[5]+" tidak dikenal")
		os.Exit(0)
	}
	if os.Args[3] == "-t" || os.Args[3] == "--table" {
		if len(os.Args[3:]) == 1 {
			fmt.Println(string("\033[31m"), "empty table name")
			return
		}
		migration.DirMigration = dirTableMigration
		migration.TableMigration = os.Args[4]
		helper.CreateMigrationFile(&migration)
		helper.ReadeMiggrationFileInFolder()
		helper.ReadeDownFileInFolder()
		helper.ReadeDropFileInFolder()
	} else {
		fmt.Println(string("\033[31m"), "argumen "+os.Args[3]+" tidak dikenal")
		fmt.Println(string("\033[00m"), "please run: gomig create --help")
		os.Exit(0)
	}
}

func CreateSeeder() {
	var seed helper.Seeder
	if len(os.Args[3:]) == 0 {
		fmt.Println(string("\033[31m"), "gomig say: empty argumen")
		fmt.Println(string("\033[00m"), "please run: gomig create --help")
		return
	}

	if os.Args[3] == "--help" || os.Args[3] == "-h" || os.Args[3] == "-n" || os.Args[3] == "--name" || os.Args[3] == "-t" || os.Args[3] == "--table" {
		if os.Args[3] == "--help" || os.Args[3] == "-h" {
			fmt.Println("helper")
			return
		}

		for i := 0; i < len(os.Args[3:]); i++ {
			if os.Args[i+3] == "-n" || os.Args[i+3] == "--name" {
				if len(os.Args[i+3:]) < 2 {
					fmt.Println("empty file nama")
					return
				}
				seed.Filename = os.Args[i+4] // set nama file
			}
			if os.Args[i+3] == "-t" || os.Args[i+3] == "--table" {
				if len(os.Args[i+3:]) < 2 {
					fmt.Println("empty target tabel name")
					return
				}
				seed.TableName = os.Args[i+4] // set nama file
			}
		}
		seed.DirSeeder = dirSeed
		if seed.Filename == "" {
			fmt.Println("empty file name")
			return
		}
		if seed.TableName == "" {
			fmt.Println("empty target tabel name")
			return
		}

		var file, _ = os.OpenFile(helper.MyRootDir(), os.O_RDWR, 0644)
		defer file.Close()
		// fmt.Println(file)

		files, err := ioutil.ReadDir("migration/database/seed")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
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
		helper.ReadeSeederFileInFolder()
		helper.ReadeDownFileInFolder()
		helper.ReadeDropFileInFolder()
	} else {
		fmt.Println("argumen not found")
		os.Exit(0)
	}
}
