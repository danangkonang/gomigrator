package helper

import (
	"fmt"
	"os"
)

func CreateTemplateFile(rootDir, tmplt string) {
	path := tmplt + "/template.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package templates\n\n"

		writeText += "var UsageTemplate = `{{.Name}} Version {{.Version}}\n\n"
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
		writeText += "`\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
