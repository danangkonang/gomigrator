package helper

import (
	"fmt"
	"os"
)

func CreatePrintTemplate(rootDir, targetDir string) {
	path := targetDir + "/create_template.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package helper\n\n"
		writeText += "import (\n"
		writeText += `	"os"`
		writeText += "\n"
		writeText += `	"text/template"`
		writeText += "\n\n"
		writeText += `	"github.com/danangkonang/` + rootDir + `/migration/app/templates"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type Inventory struct {"
		writeText += "\n"
		writeText += "	Name    string"
		writeText += "\n"
		writeText += "	Version string"
		writeText += "\n"
		writeText += "}"
		writeText += "\n\n"

		// type Inventory struct {
		// 	Name    string
		// 	Version string
		// }

		// function
		writeText += "func PrintHelper() {\n"
		writeText += `	data := Inventory{"Danang", "0.0.1"}`
		writeText += "\n"
		writeText += `	tmpl, err := template.New("test").Parse(templates.UsageTemplate)`
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

		writeText += "func PrintVersion() {\n"
		writeText += `	data := Inventory{"Danang", "0.0.1"}`
		writeText += "\n"
		writeText += `	tmpl, err := template.New("test").Parse(templates.UsageTemplate)`
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

		file.WriteString(writeText)
		defer file.Close()
	}
}
