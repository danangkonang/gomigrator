package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var pathDrop = "migration/app/execusion/drop_table.go"

func ReadeDownFileInFolder() {
	deleteFiles()
	createFiles()
	writeFiles()
}

func writeFiles() {
	var file, err = os.OpenFile(pathDrop, os.O_RDWR, 0644)
	if isErrors(err) {
		return
	}
	defer file.Close()

	files, err := ioutil.ReadDir("migration/database/migration")
	if err != nil {
		log.Fatal(err)
	}

	writeText := "package execusion\n\n"
	if len(files) != 0 {
		writeText += "import (\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"io/ioutil"`
		writeText += "\n"
		writeText += `	"os"`
		writeText += "\n"
		writeText += `	"strings"`
		writeText += "\n\n"
		writeText += `	"github.com/danangkonang/` + MyRootDir() + `/migration/app/config"`
		writeText += "\n"
		writeText += ")\n\n"
	}

	// writeText += "type Tables struct {\n"
	// writeText += "	NameTable []string\n"
	// writeText += "}\n\n"

	// function
	writeText += "func DropTables(tb *Tables) {\n"
	if len(files) != 0 {
		writeText += `	files, err := ioutil.ReadDir("migration/database/migration")`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err)\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	db := config.Connect()\n"

		writeText += "	if len(tb.NameTable) > 0 {\n"

		writeText += "		for _, ntb := range tb.NameTable {\n"
		writeText += `			query := "DROP TABLE IF EXISTS " + ntb + ";"`
		writeText += "\n"
		writeText += "			_, err := db.Exec(query)\n"
		writeText += "			if err != nil {\n"
		writeText += "				fmt.Println(err)\n"
		writeText += "				os.Exit(0)\n"
		writeText += "			}\n"
		writeText += `			fmt.Println("success DROP TABLE " + ntb)`
		writeText += "\n"
		writeText += "		}\n"

		writeText += "	} else {\n"

		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			list := strings.Split(filename, "_migration_")`
		writeText += "\n"
		writeText += "			name := list[0]\n"
		writeText += `			if name != "0.core_type_migration.go" {`
		writeText += "\n"
		writeText += `				query := "DROP TABLE IF EXISTS " + name + ";"`
		writeText += "\n"
		writeText += "				_, err := db.Exec(query)\n"
		writeText += "				if err != nil {\n"
		writeText += "					fmt.Println(err)\n"
		writeText += "					os.Exit(0)\n"
		writeText += "				}\n"
		writeText += `				fmt.Println("success DROP TABLE " + name)`
		writeText += "\n"
		writeText += "			}\n"
		writeText += "		}\n" // end for files

		writeText += "	}\n" // end else
	} // end len(files)

	writeText += "}\n"
	// writeText += "\n"

	_, err = file.WriteString(writeText)
	if isErrors(err) {
		return
	}
	// save changes
	err = file.Sync()
	if isErrors(err) {
		return
	}
}

func createFiles() {
	var _, err = os.Stat(pathDrop)
	if os.IsNotExist(err) {
		var file, err = os.Create(pathDrop)
		if isErrors(err) {
			return
		}
		defer file.Close()
	}
}

func deleteFiles() {
	var _, err = os.Stat(pathDrop)
	if err == nil {
		var err = os.Remove(pathDrop)
		if isErrors(err) {
			return
		}
	}
}

func isErrors(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
