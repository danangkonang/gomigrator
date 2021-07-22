package command

import (
	"fmt"
	"os"

	"github.com/danangkonang/migration-go-cli/app/helper"
)

func CreateMigrationTableV2() {
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
		migration.DirMigration = DirMigration
		migration.TableMigration = os.Args[4]
		helper.CreateMigrationFileV2(&migration)
		// helper.ReadeMiggrationFileInFolder()
		// helper.ReadeDownFileInFolder()
		// helper.ReadeDropFileInFolder()
	} else {
		fmt.Println(string("\033[31m"), "argumen "+os.Args[3]+" tidak dikenal")
		fmt.Println(string("\033[00m"), "please run: gomig create --help")
		os.Exit(0)
	}
}
