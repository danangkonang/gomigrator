package helper

import (
	"os"
	"text/template"

	"github.com/danangkonang/migrasion-go-cli/templates"
)

type Inventory struct {
	Name    string
	Version string
}

type ERROR struct {
	Message string
}

func PrintHelper() {
	data := Inventory{"gomig", "0.0.2"}
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
