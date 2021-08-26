package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/helper"
)

func CreateSeederTableV2() {
	var seeder helper.Seeder
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
				seeder.Filename = os.Args[i+4] // set nama file
			}
			if os.Args[i+3] == "-t" || os.Args[i+3] == "--table" {
				if len(os.Args[i+3:]) < 2 {
					fmt.Println("empty target tabel name")
					return
				}
				seeder.TableName = os.Args[i+4] // set nama file
			}
		}
		seeder.DirSeeder = DirSeeder
		if seeder.Filename == "" {
			fmt.Println("empty file name")
			return
		}
		if seeder.TableName == "" {
			fmt.Println("empty target tabel name")
			return
		}

		var file, _ = os.OpenFile(helper.MyRootDir(), os.O_RDWR, 0644)
		defer file.Close()
		// fmt.Println(file)

		files, err := ioutil.ReadDir(DirSeeder)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		for _, file := range files {
			filename := file.Name()
			list := strings.Split(filename, "_seeder_")
			if list[0] != "0.go" {
				fl_name := strings.Split(list[1], ".go")
				if fl_name[0] == seeder.Filename {
					fmt.Println(string("\033[31m"), "seeder file name musth be unique")
					os.Exit(0)
				}
			}
			// name := list[0]
		}

		helper.CreateSeedFileV2(&seeder)

		// helper.ReadeSeederFileInFolder()
		// helper.ReadeDownFileInFolder()
		// helper.ReadeDropFileInFolder()
	} else {
		fmt.Println("argumen not found")
		os.Exit(0)
	}
}
