package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var path = "migration/app/execusion/reset_table.go"

func ReadeDropFileInFolder() {
	deleteFile()
	createFile()
	writeFile()
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
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

	// function
	writeText += "func ResetTables(tb *Tables) {\n"
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
		writeText += `			query := "TRUNCATE " + ntb + ";"`
		writeText += "\n"
		writeText += "			_, err := db.Exec(query)\n"
		writeText += "			if err != nil {\n"
		writeText += "				fmt.Println(err)\n"
		writeText += "				os.Exit(0)\n"
		writeText += "			}\n"
		writeText += `			fmt.Println("success delete row")`
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
		writeText += `				query := "TRUNCATE " + name + ";"`
		writeText += "\n"
		writeText += "				_, err := db.Exec(query)\n"
		writeText += "				if err != nil {\n"
		writeText += "					fmt.Println(err)\n"
		writeText += "					os.Exit(0)\n"
		writeText += "				}\n"
		writeText += `				fmt.Println("success delete row")`
		writeText += "\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "	}\n" // end else
	}
	writeText += "}\n"
	// writeText += "\n"

	_, err = file.WriteString(writeText)
	if isError(err) {
		return
	}
	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}
}

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}

func deleteFile() {
	var _, err = os.Stat(path)
	if err == nil {
		var err = os.Remove(path)
		if isError(err) {
			return
		}
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
