// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// var dir = "database"
// var path = dir + "/" + getTime() + "-test.go"

// func isError(err error) bool {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return (err != nil)
// }

// func getTime() string {
// 	t := time.Now()
// 	waktu := t.Format("2006-01-02-150405")
// 	return waktu
// }

// func createFile() {
// 	// deteksi apakah file sudah ada
// 	_, err := os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		file.WriteString("package database")
// 		defer file.Close()
// 	}
// 	fmt.Println("file berhasil dibuat", path)
// }

// func createFolder() {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		os.Mkdir(dir, 0700)
// 		fmt.Println(dir + " telah dibuat")
// 	}
// }

// func main() {
// 	createFolder()
// 	createFile()
// 	fmt.Println(getTime())
// }

package main

import (
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/cmd"
	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var dir = "database"

// var path = dir + "/" + getTime() + "-test.go"

// func isError(err error) bool {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return (err != nil)
// }

// func cli() {
// 	name := flag.String("migrate", "", "make file migration 'konang -migrate users'")
// 	// seed := flag.String("migrate", "", "make file migration 'konang -migrate users'")
// 	flag.Parse()
// 	if len(*name) == 0 {
// 		// os.Exit(1)
// 		helper.PrintHelper()
// 		return
// 	}
// 	createFolder(name)
// }

// func createFolder(name *string) {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		os.Mkdir(dir, 0700)
// 		fmt.Println(dir + " telah dibuat")
// 	}
// 	createFile(name)
// }

// func createFile(name *string) {
// 	path := dir + "/" + getTime() + "-" + *name + ".go"
// 	// deteksi apakah file sudah ada
// 	_, err := os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		file.WriteString("package database\n")
// 		defer file.Close()
// 	}
// 	fmt.Println("file berhasil dibuat", path)
// }

// func getTime() string {
// 	t := time.Now()
// 	waktu := t.Format("2006-01-02-150405")
// 	return waktu
// }

func main() {
	arrCmd := os.Args[1:]
	if len(arrCmd) == 0 {
		helper.PrintHelper()
		return
	}
	runCmd()
}

func runCmd() {
	usrCmd := os.Args[1]

	switch usrCmd {
	case "-h":
		helper.PrintHelper()
		break
	case "--help":
		helper.PrintHelper()
		break
	case "-v":
		helper.PrintVersion()
		break
	case "--version":
		helper.PrintVersion()
		break
	case "init":
		cmd.Initial()
		break
	case "danang":
		cmd.CmdType(usrCmd)
		break
	case "create":
		cmd.TypeIs()
		break
	case "run":
		cmd.TypeIs()
		break
	case "back":
		cmd.TypeIs()
		break
	default:
		helper.PrintHelper()
	}
}

// type Inventory struct {
// 	Name    string
// 	Version string
// }

// type Comman struct {
// 	Cmd []string
// }

// func hepl() {
// 	data := Inventory{"Danang", "0.0.1"}
// 	tmpl, err := template.New("test").Parse(templates.UsageTemplate)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = tmpl.Execute(os.Stdout, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func version() {
// 	data := Inventory{"Danang", "0.0.1"}
// 	tmpl, err := template.New("test").Parse(templates.VersionTemplate)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = tmpl.Execute(os.Stdout, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// var initdir = "config"

// func initial() {
// 	if _, err := os.Stat(initdir); os.IsNotExist(err) {
// 		os.Mkdir(initdir, 0700)
// 		fmt.Println(initdir + " telah dibuat")
// 	}
// 	makeEnvFile()
// }

// func makeEnvFile() {
// 	path := ".env"
// 	// deteksi apakah file sudah ada
// 	_, err := os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		file.WriteString(`DBNAME=root\n
// DBPASSWORD=root\n
// DBUSER=postgres\n
// DBTYPE=postgres\n
// DBHOST=localhost\n
// DBPORT=5432\n
// PORT=9000\n`)
// 		defer file.Close()
// 	}
// 	fmt.Println("file berhasil dibuat", path)
// 	makeConfigFile()
// }

// func makeConfigFile() {
// 	path := initdir + "/config.go"
// 	// deteksi apakah file sudah ada
// 	_, err := os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		file.WriteString("package config\n")
// 		defer file.Close()
// 	}
// 	fmt.Println("file berhasil dibuat", path)
// }
