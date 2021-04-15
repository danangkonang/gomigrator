package execusion

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var pathSeed = "migration/app/execusion/runing_seeder.go"

func ReadeSeederFileInFolder() {
	rootDir := helper.MyRootDir()
	deleteFileSeed()
	createFileSeed()
	writeFileSeed(rootDir)
}

func createFileSeed() {
	// detect if file exists
	var _, err = os.Stat(pathSeed)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(pathSeed)
		if isErrorSeed(err) {
			return
		}
		defer file.Close()
	}
	// fmt.Println("==> done creating file", pathSeed)
}

func writeFileSeed(thisDir string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(pathSeed, os.O_RDWR, 0644)
	if isErrorSeed(err) {
		return
	}
	defer file.Close()

	files, err := ioutil.ReadDir("migration/database/seed")
	// fmt.Println("panjang", len(files))
	if err != nil {
		log.Fatal(err)
	}

	writeText := "package execusion\n\n"
	if len(files) != 0 {
		writeText += "import (\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/database/seed"`
		writeText += "\n"
		writeText += ")\n\n"
	}

	// function
	writeText += "func RuningSeeder() {\n"

	for _, file := range files {
		filename := file.Name()
		list := strings.Split(filename, "_seed_")
		name := list[0]
		writeText += "\tseed." + strings.Title(name+"()\n")
	}

	writeText += "}\n"
	// write some text line-by-line to file
	_, err = file.WriteString(writeText)
	if isErrorSeed(err) {
		return
	}
	// save changes
	err = file.Sync()
	if isErrorSeed(err) {
		return
	}
	// fmt.Println("==> done writing to file")
}

func deleteFileSeed() {
	var _, err = os.Stat(pathSeed)
	if err == nil {
		var err = os.Remove(pathSeed)
		if isErrorSeed(err) {
			return
		}
	}
	// fmt.Println("==> done deleting file")
}

func isErrorSeed(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
