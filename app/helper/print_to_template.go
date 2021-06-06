package helper

import (
	"os"
	"text/template"

	"github.com/danangkonang/migration-go-cli/templates"
)

type Inventory struct {
	Name    string
	Version string
}

type ERROR struct {
	Message string
}

type Helper struct {
	AppName       string
	AppVersion    string
	Cmd           string
	Version       string
	Helper        string
	Table         string
	Name          string
	Migration     string
	Seeder        string
	Init          string
	Reset         string
	Drop          string
	Create        string
	Run           string
	VersionDesc   string
	HelperDesc    string
	TableDesc     string
	NameDesc      string
	MigrationDesc string
	SeederDesc    string
	InitDesc      string
	ResetDesc     string
	DropDesc      string
	CreateDesc    string
	RunDesc       string
}

func PrintHelper() {
	// data := Inventory{"gomig", "0.0.2"}
	data := &Helper{

		Cmd:           "",
		AppVersion:    "0.0.2",
		AppName:       "gomig",
		Version:       "-v, --version",
		VersionDesc:   "print version cli",
		Helper:        "-h, --help",
		HelperDesc:    "output usage information",
		Table:         "-t, table",
		TableDesc:     "create table name for migration",
		Name:          "-n, name",
		NameDesc:      "create file name seeder or migration",
		Migration:     "migration",
		MigrationDesc: "create migration file",
		Seeder:        "seeder",
		SeederDesc:    "create seeder file",
		Init:          "init",
		InitDesc:      "init project",
		Reset:         "reset",
		ResetDesc:     "delete seeder",
		Drop:          "drop",
		DropDesc:      "delete tables",
		Create:        "create",
		CreateDesc:    "create migration or seeder [table name]",
		Run:           "run",
		RunDesc:       "running migration or seeder",
	}
	tmpl, err := template.New("helper").Parse(templates.HelperTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func PrintHelperCreate() {
	data := Inventory{"gomig", "0.0.2"}
	tmpl, err := template.New("helper").Parse(templates.HelperCreateTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func PrintHelperRun() {
	data := Inventory{"gomig", "0.0.2"}
	tmpl, err := template.New("helper").Parse(templates.HelperRunTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func MultyPrintHelper(comand string) {
	var data = &Helper{
		AppName:       "gomig",
		Version:       "0.0.2",
		Cmd:           comand,
		Table:         "-t, table",
		TableDesc:     "create table name for migration",
		Name:          "-n, name",
		NameDesc:      "create file name seeder or migration",
		Migration:     "migration",
		MigrationDesc: "create migration file",
		Seeder:        "seeder",
		SeederDesc:    "create seeder file",
	}

	tmpl, err := template.New("helper").Parse(templates.HelperRunTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func PrintVersion() {
	data := Inventory{"gomig", "0.0.2"}
	tmpl, err := template.New("version").Parse(templates.VersionTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func ErrorCommand(userFlag string) {
	data := ERROR{userFlag}
	tmpl, err := template.New("test").Parse(templates.ErrorTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
