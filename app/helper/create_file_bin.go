package helper

import (
	"fmt"
	"os"
)

func CreateBinFileNew(thisDir, dirMigration string) {
	path := dirMigration + "/bin.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package main\n\n"
		writeText += "import (\n"
		writeText += `	"flag"`
		writeText += "\n"
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
		writeText += `	"github.com/joho/godotenv"`
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
		writeText += `	green           = "\033[32m"`
		writeText += "\n"
		writeText += `	reset           = "\033[0m"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "func init() {\n"
		writeText += "	godotenv.Load()\n"
		writeText += "}\n\n"

		/*
		* Main
		 */
		writeText += "func main() {\n"
		writeText += "	var t Tables\n"
		writeText += `	up := flag.NewFlagSet("up", flag.ExitOnError)`
		writeText += "\n"
		writeText += `	upMigration := flag.NewFlagSet("migration", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	upMigration.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += `	upSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	upSeeder.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n\n"

		writeText += `	down := flag.NewFlagSet("down", flag.ExitOnError)`
		writeText += "\n"
		writeText += `	downMigration := flag.NewFlagSet("migration", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	downMigration.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += `	downSeeder := flag.NewFlagSet("seeder", flag.ContinueOnError)`
		writeText += "\n"
		writeText += `	downSeeder.Func("tables", "list file name", func(s string) error {`
		writeText += "\n"
		writeText += "		t.NameTable = strings.Fields(s)\n"
		writeText += "		return nil\n"
		writeText += "	})\n"
		writeText += "	if len(os.Args) < 2 {\n"
		writeText += `		fmt.Println("helper")`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	switch os.Args[1] {\n"
		writeText += `	case "up":`
		writeText += "\n"
		writeText += "		up.Parse(os.Args[2:])\n"
		writeText += "		upHandle(os.Args, upMigration, upSeeder, &t)\n"
		writeText += `	case "down":`
		writeText += "\n"
		writeText += "		down.Parse(os.Args[2:])\n"
		writeText += "		downHandle(os.Args, upMigration, upSeeder, &t)\n"
		writeText += "	default:\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func downHandle(argument []string, upMigration, upSeeder *flag.FlagSet, tb *Tables) {\n"
		writeText += "	switch argument[2] {\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		upMigration.Parse(os.Args[3:])\n"
		writeText += "		migrationDown(tb)\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		upSeeder.Parse(os.Args[3:])\n"
		writeText += "		seederDown(tb)\n"
		writeText += "	default:\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func migrationDown(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(MigrationFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println("help me")`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		writeText += `		list := strings.Split(migrate, "_migration_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += `		query := "DROP TABLE IF EXISTS " + tb_name[0] + " CASCADE;"`
		writeText += "\n"
		writeText += "		_, err := migration.Connection().Db.Exec(query)\n"
		writeText += "		if err != nil {\n"
		writeText += "			fmt.Println(err)\n"
		writeText += "			os.Exit(0)\n"
		writeText += "		}\n"
		writeText += `		fmt.Println(string(green), "success", string(reset), "DROP ", migrate)`
		writeText += "\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func seederDown(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(SeederFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println("help me")`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		writeText += `		list := strings.Split(migrate, "_seeder_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += `		query := "TRUNCATE " + tb_name[0] + " RESTART IDENTITY;"`
		writeText += "\n"
		writeText += "		_, err := migration.Connection().Db.Exec(query)\n"
		writeText += "		if err != nil {\n"
		writeText += "			fmt.Println(err.Error())\n"
		writeText += "			os.Exit(0)\n"
		writeText += "		}\n"
		writeText += `		fmt.Println(string(green), "success", string(reset), "remove data on TABLE ", migrate)`
		writeText += "\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func upHandle(argument []string, upMigration, upSeeder *flag.FlagSet, tb *Tables) {\n"
		writeText += "	switch argument[2] {\n"
		writeText += `	case "migration":`
		writeText += "\n"
		writeText += "		upMigration.Parse(os.Args[3:])\n"
		writeText += "		migrationUp(tb)\n"
		writeText += `	case "seeder":`
		writeText += "\n"
		writeText += "		upSeeder.Parse(os.Args[3:])\n"
		writeText += "		seederUp(tb)\n"
		writeText += "	default:\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func migrationUp(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(MigrationFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println("help me")`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	m := migration.Migration{}\n"
		writeText += "	for _, migrate := range t.NameTable {\n"
		writeText += `		list := strings.Split(migrate, "_migration_")`
		writeText += "\n"
		writeText += `		tb_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += "		meth := reflect.ValueOf(&m).MethodByName(strings.Title(tb_name[0]))\n"
		writeText += "		meth.Call(nil)\n"
		writeText += "	}\n"
		writeText += "}\n\n"

		writeText += "func seederUp(t *Tables) {\n"
		writeText += "	files, err := ioutil.ReadDir(SeederFolder)\n"
		writeText += "	if err != nil {\n"
		writeText += `		fmt.Println("help me")`
		writeText += "\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	if len(t.NameTable) == 0 {\n"
		writeText += "		newFile := []string{}\n"
		writeText += "		for _, file := range files {\n"
		writeText += "			filename := file.Name()\n"
		writeText += `			if filename != "0.go" {`
		writeText += "\n"
		writeText += "				newFile = append(newFile, filename)\n"
		writeText += "			}\n"
		writeText += "		}\n"
		writeText += "		t.NameTable = newFile\n"
		writeText += "	}\n"
		writeText += "	s := seeder.Seeder{}\n"
		writeText += "	for _, seed := range t.NameTable {\n"
		writeText += `		list := strings.Split(seed, "_seeder_")`
		writeText += "\n"
		writeText += `		func_name := strings.Split(list[1], ".go")`
		writeText += "\n"
		writeText += "		meth := reflect.ValueOf(&s).MethodByName(strings.Title(func_name[0]))\n"
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

		file.WriteString(writeText)
		defer file.Close()
	}
}
