package execusion

import (
	"fmt"
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var path_seed = "migration/app/execusion/runing_seeder.go"

func ReadeSeederFileInFolder() {
	rootDir := helper.MyRootDir()
	deleteFileSeed()
	createFileSeed()
	writeFileSeed(rootDir)
}

func createFileSeed() {
	// detect if file exists
	var _, err = os.Stat(path_seed)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path_seed)
		if isErrorSeed(err) {
			return
		}
		defer file.Close()
	}
	// fmt.Println("==> done creating file", path_seed)
}

func writeFileSeed(thisDir string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path_seed, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// files, err := ioutil.ReadDir("migration/database/migration")
	// // fmt.Println("panjang", len(files))
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	writeText += `	"github.com/danangkonang/` + thisDir + `/migration/database/seed"`
	writeText += "\n"
	writeText += ")\n\n"
	// }

	// function
	writeText += "func RuningSeeder(tbl *Tables) {\n"

	// for _, file := range files {
	// 	filename := file.Name()
	// 	list := strings.Split(filename, "_migration_")
	// 	name := list[0]
	// 	writeText += "\tmigration." + strings.Title(name+"()\n")
	// }

	// if len(files) != 0 {

	writeText += `	files, err := ioutil.ReadDir("migration/database/seed")`
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
	writeText += `			list := strings.Split(filename, "_seeder_")`
	writeText += "\n"
	writeText += "			name := list[0]"
	writeText += "\n"
	writeText += `			if name != "0.core_type_seed.go" {`
	writeText += "\n"
	writeText += "				newFile = append(newFile, name)"
	writeText += "\n"
	writeText += "			}"
	writeText += "\n"
	writeText += "		}"
	writeText += "\n"
	writeText += "		tbl.NameTable = newFile"
	writeText += "\n"
	writeText += "	}"
	writeText += "\n"

	writeText += "	s := seed.MySeed{}"
	writeText += "\n"
	writeText += "	for _, data_seeder := range tbl.NameTable {"
	writeText += "\n"
	writeText += "		meth := reflect.ValueOf(s).MethodByName(strings.Title(data_seeder))"
	writeText += "\n"
	writeText += "		meth.Call(nil)"
	writeText += "\n"
	writeText += "	}"
	writeText += "\n"
	// }
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
}

func deleteFileSeed() {
	var _, err = os.Stat(path_seed)
	if err == nil {
		var err = os.Remove(path_seed)
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
