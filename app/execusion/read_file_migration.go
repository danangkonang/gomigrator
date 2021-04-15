package execusion

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var path = "migration/app/execusion/runing_migration.go"

func ReadeMiggrationFileInFolder() {
	deleteFile()
	createFile()
	writeFile(helper.MyRootDir())
}

func createFile() {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	// fmt.Println("==> done creating file", path)
}

func writeFile(thisDir string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	files, err := ioutil.ReadDir("migration/database/migration")
	// fmt.Println("panjang", len(files))
	if err != nil {
		log.Fatal(err)
	}

	writeText := "package execusion\n\n"
	if len(files) != 0 {
		writeText += "import (\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/database/migration"`
		writeText += "\n"
		writeText += ")\n\n"
	}

	// function
	writeText += "func RuningMigration() {\n"

	for _, file := range files {
		filename := file.Name()
		list := strings.Split(filename, "_migration_")
		name := list[0]
		writeText += "\tmigration." + strings.Title(name+"()\n")
	}
	writeText += "}\n"
	// write some text line-by-line to file
	_, err = file.WriteString(writeText)
	if isError(err) {
		return
	}
	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	// fmt.Println("==> done writing to file")
}

func deleteFile() {
	// delete file
	var _, err = os.Stat(path)
	if err == nil {
		var err = os.Remove(path)
		if isError(err) {
			return
		}
	}
	// fmt.Println("==> done deleting file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
