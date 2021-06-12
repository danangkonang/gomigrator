package helper

import (
	"fmt"
	"os"
)

// import (
// 	"fmt"
// 	"os"
// )

var path_migration = "migration/app/execusion/runing_migration.go"

func ReadeMiggrationFileInFolder() {
	deleteFileRm()
	createFileRM()
	writeFileRM(MyRootDir())
}

func createFileRM() {
	// detect if file exists
	var _, err = os.Stat(path_migration)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path_migration)
		if isErrorRm(err) {
			return
		}
		defer file.Close()
	}

	// fmt.Println("==> done creating file", path_migration)
}

func writeFileRM(thisDir string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path_migration, os.O_RDWR, 0644)
	if isErrorRm(err) {
		return
	}
	defer file.Close()

	writeText := "package execusion\n\n"
	// if len(files) != 0 {
	writeText += "import (\n"
	writeText += `	"io/ioutil"`
	writeText += "\n"
	writeText += `	"os"`
	writeText += "\n"
	writeText += "\n"
	writeText += `	"reflect"`
	writeText += "\n"
	writeText += `	"strings"`
	writeText += "\n"
	writeText += "\n"
	writeText += `	"github.com/danangkonang/` + thisDir + `/migration/database/migration"`
	writeText += "\n"
	writeText += ")\n\n"

	// function
	writeText += "func RuningMigration(tbl *Tables) {\n"

	writeText += `	files, err := ioutil.ReadDir("migration/database/migration")`
	writeText += "\n"
	writeText += "	if err != nil {"
	writeText += "\n"
	writeText += "		os.Exit(0)"
	writeText += "\n"
	writeText += "	}"
	writeText += "\n"

	writeText += "	if len(tbl.NameTable) == 0 {"
	writeText += "\n"
	writeText += "		newFile := []string{}"
	writeText += "\n"
	writeText += "		for _, file := range files {"
	writeText += "\n"
	writeText += "			filename := file.Name()"
	writeText += "\n"
	writeText += `			list := strings.Split(filename, "_migration_")`
	writeText += "\n"
	writeText += "			name := list[1]"
	writeText += "\n"
	writeText += `			tb_name := strings.Split(name, ".go")`
	writeText += "\n"
	writeText += `			if name != "0.core_type_migration.go" {`
	writeText += "\n"
	writeText += "				newFile = append(newFile, tb_name[0])"
	writeText += "\n"
	writeText += "			}"
	writeText += "\n"
	writeText += "		}"
	writeText += "\n"
	writeText += "		tbl.NameTable = newFile"
	writeText += "\n"
	writeText += "	}"
	writeText += "\n"

	writeText += "	m := migration.MyMigration{}"
	writeText += "\n"
	writeText += "	for _, migrate := range tbl.NameTable {"
	writeText += "\n"
	writeText += "		meth := reflect.ValueOf(m).MethodByName(strings.Title(migrate))"
	writeText += "\n"
	writeText += "		meth.Call(nil)"
	writeText += "\n"
	writeText += "	}"
	writeText += "\n"
	// }
	writeText += "}\n"
	// write some text line-by-line to file
	_, err = file.WriteString(writeText)
	if isErrorRm(err) {
		return
	}
	// save changes
	err = file.Sync()
	if isErrorRm(err) {
		return
	}

	// fmt.Println("==> done writing to file")
}

func deleteFileRm() {
	// delete file
	var _, err = os.Stat(path_migration)
	if err == nil {
		var err = os.Remove(path_migration)
		if isErrorRm(err) {
			return
		}
	}
	// fmt.Println("==> done deleting file")
}

func isErrorRm(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
