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

func PrintHelper() {
	data := Inventory{"Danang", "0.0.1"}
	tmpl, err := template.New("test").Parse(templates.UsageTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func PrintVersion() {
	data := Inventory{"Danang", "0.0.1"}
	tmpl, err := template.New("test").Parse(templates.VersionTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
