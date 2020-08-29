package helper

import (
	"fmt"
	"os"
)

var dirConfig = "config"

func CreateConfigFile() {
	path := dirConfig + "/config.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.WriteString("package config\n")
		defer file.Close()
	}
}
