package helper

import (
	"fmt"
	"os"
)

func CreateBinFile(thisDir, dirMigration string) {
	path := dirMigration + "/bin.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package main\n\n"
		writeText += "import (\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"io/ioutil"`
		writeText += "\n"
		writeText += `	"os"`
		writeText += "\n"
		writeText += `	"reflect"`
		writeText += "\n"
		writeText += `	"strings"`
		writeText += "\n"
		writeText += `	"text/template"`
		writeText += "\n\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/db/migration"`
		writeText += "\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/db/seeder"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type Tables struct {\n"
		writeText += "	NameTable []string\n"
		writeText += "}\n\n"

		writeText += "var (\n"
		writeText += `	MigrationFolder = "db/migration"`
		writeText += "\n"
		writeText += `	SeederFolder    = "db/seeder"`
		writeText += "\n"
		writeText += ")\n"
		// function
		writeText += "func main() {\n"
		writeText += "	if len(os.Args[1:]) == 0 {\n"
		writeText += "		PrintHelper()\n"
		writeText += "		return\n"
		writeText += "	}\n"
		writeText += "	runCmd()\n"
		writeText += "}\n"
		writeText += "\n"
		writeText += "func runCmd() {\n"
		writeText += "	var t Tables\n"
		writeText += "	switch os.Args[1] {\n"
		writeText += `	case "run":`
		writeText += "\n"
		writeText += "		migrationOrSeeder()"
		writeText += "\n"
		writeText += `	case "reset":`
		writeText += "\n"
		writeText += `		if os.Args[2] != "" {`
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[2], ",")`
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "		ResetTables(&t)"
		writeText += "\n"
		writeText += `	case "drop":`
		writeText += "\n"
		writeText += `		if os.Args[2] != "" {`
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[2], ",")`
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "		DropTables(&t)"
		writeText += "\n"
		writeText += "	default:\n"
		writeText += "		PrintHelper()\n"
		writeText += "	}\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func migrationOrSeeder() {"
		writeText += "\n"
		writeText += "	var t Tables\n"
		writeText += "	switch os.Args[2] {"
		writeText += "\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "\n"
		writeText += `		if os.Args[3] == "" {`
		writeText += "\n"
		writeText += "			RuningMigration(&t)"
		writeText += "\n"
		writeText += "		} else {"
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[3], ",")`
		writeText += "\n"
		writeText += "			RuningMigration(&t)"
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "\n"
		writeText += `		if os.Args[3] == "" {`
		writeText += "\n"
		writeText += "			RuningSeeder(&t)"
		writeText += "\n"
		writeText += "		} else {"
		writeText += "\n"
		writeText += `			t.NameTable = strings.Split(os.Args[3], ",")`
		writeText += "\n"
		writeText += "			RuningSeeder(&t)"
		writeText += "\n"
		writeText += "		}"
		writeText += "\n"
		writeText += "	default:"
		writeText += "\n"
		writeText += "		PrintHelper()"
		writeText += "\n"
		writeText += "	}"
		writeText += "\n"
		writeText += "}"
		writeText += "\n\n"

		/*
		* RuningSeeder
		 */
		writeText += "func RuningMigration(tbl *Tables) {\n"
		writeText += `	files, err := ioutil.ReadDir(MigrationFolder)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(tbl.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			list := strings.Split(filename, "_migration_")`
		writeText += "\n"
		writeText += `			if list[0] != "0.go" {`
		writeText += "\n"
		writeText += "				name := list[1]\n"
		writeText += `				tb_name := strings.Split(name, ".go")`
		writeText += "\n"
		writeText += "				newFile = append(newFile, tb_name[0])\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		tbl.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	m := migration.Migration{}\n"
		writeText += "	for _, migrate := range tbl.NameTable {\n"
		writeText += "		meth := reflect.ValueOf(&m).MethodByName(strings.Title(migrate))\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		/*
		* RuningSeeder
		 */
		writeText += "func RuningSeeder(tbl *Tables) {\n"
		writeText += `	files, err := ioutil.ReadDir(SeederFolder)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(tbl.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			list := strings.Split(filename, "_seeder_")`
		writeText += "\n"
		writeText += `			if list[0] != "0.go" {`
		writeText += "\n"
		writeText += "				name := list[1]\n"
		writeText += `				tb_name := strings.Split(name, ".go")`
		writeText += "\n"
		writeText += "				newFile = append(newFile, tb_name[0])\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		tbl.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	m := seeder.Seeder{}\n"
		writeText += "	for _, migrate := range tbl.NameTable {\n"
		writeText += "		meth := reflect.ValueOf(&m).MethodByName(strings.Title(migrate))\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		/*
		* Gomig
		 */
		writeText += "type Gomig struct {"
		writeText += "\n"
		writeText += "	Name    string"
		writeText += "\n"
		writeText += "	Version string"
		writeText += "\n"
		writeText += "}"
		writeText += "\n\n"

		/*
		* template
		 */
		writeText += "var MyTemplate = `{{.Name}} Version {{.Version}}\n\n"
		writeText += `Usage: {{.Name}} [command] [options]`
		writeText += "\n\n"
		writeText += "Options:\n"
		writeText += "	-v, --version                       output the version number\n"
		writeText += "	-h, --help                          output usage information\n\n"
		writeText += "Commands:\n"
		writeText += "	- migrate\n"
		writeText += "	- run\n"
		writeText += "	- down\n\n"
		writeText += "{{- /* end */ -}}\n"
		writeText += `{{- "" }}`
		writeText += "`\n\n"
		writeText += "var VersionTemplate = `{{.Name}} Version {{.Version}}\n"
		writeText += "{{- /* end */ -}}\n"
		writeText += `{{- "" }}`
		writeText += "\n"
		writeText += "`\n\n"

		/*
		* PrintHelper
		 */
		writeText += "func PrintHelper() {\n"
		writeText += `	data := Gomig{"Danang", "0.0.5"}`
		writeText += "\n"
		writeText += `	tmpl, err := template.New("test").Parse(MyTemplate)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "	err = tmpl.Execute(os.Stdout, data)\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "}\n"
		writeText += "\n"

		/*
		* PrintVersion
		 */
		writeText += "func PrintVersion() {\n"
		writeText += `	data := Gomig{"Danang", "0.0.5"}`
		writeText += "\n"
		writeText += `	tmpl, err := template.New("test").Parse(MyTemplate)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "	err = tmpl.Execute(os.Stdout, data)\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "}\n"
		writeText += "\n"

		/*
		* DropTables
		 */
		writeText += "func DropTables(tb *Tables) {\n"
		writeText += `	files, err := ioutil.ReadDir(MigrationFolder)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err)\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(tb.NameTable) > 0 {\n"
		writeText += "		for _, ntb := range tb.NameTable {\n"
		writeText += `			query := "DROP TABLE IF EXISTS " + ntb + " CASCADE;"`
		writeText += "\n"
		writeText += "			_, err := migration.Connection().Db.Exec(query)\n"
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
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += `				names := strings.Split(filename, "_migration_")`
		writeText += "\n"
		writeText += `				tb_name := strings.Split(names[1], ".go")`
		writeText += "\n"
		writeText += `				query := "DROP TABLE IF EXISTS " + tb_name[0] + " CASCADE;"`
		writeText += "\n"
		writeText += "				_, err := migration.Connection().Db.Exec(query)\n"
		writeText += "				if err != nil {\n"
		writeText += "					fmt.Println(err)\n"
		writeText += "					os.Exit(0)\n"
		writeText += "				}\n"
		writeText += `				fmt.Println("success DROP TABLE " + tb_name[0])`
		writeText += "\n"
		writeText += "			}\n"
		writeText += "		}\n" // end for files
		writeText += "	}\n"  // end else
		writeText += "}\n\n"

		/*
		* ResetTables
		*
		 */
		writeText += "func ResetTables(tb *Tables) {\n"
		writeText += `	files, err := ioutil.ReadDir(MigrationFolder)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err)\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(tb.NameTable) > 0 {\n"
		writeText += "		for _, ntb := range tb.NameTable {\n"
		writeText += `			query := "TRUNCATE " + ntb + ";"`
		writeText += "\n"
		writeText += "			_, err := migration.Connection().Db.Exec(query)\n"
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
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += `				names := strings.Split(filename, "_migration_")`
		writeText += "\n"
		writeText += `				tb_name := strings.Split(names[1], ".go")`
		writeText += "\n"
		writeText += `				query := "TRUNCATE " + tb_name[0] + ";"`
		writeText += "\n"
		writeText += "				_, err := migration.Connection().Db.Exec(query)\n"
		writeText += "				if err != nil {\n"
		writeText += "					fmt.Println(err)\n"
		writeText += "					os.Exit(0)\n"
		writeText += "				}\n"
		writeText += `				fmt.Println("success delete row")`
		writeText += "\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "	}\n" // end else
		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
