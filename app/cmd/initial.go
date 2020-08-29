package cmd

import (
	"fmt"
	"os"
)

var dirConfig = "config"

func Initial() {
	if _, err := os.Stat(dirConfig); os.IsNotExist(err) {
		os.Mkdir(dirConfig, 0700)
		fmt.Println(dirConfig + " telah dibuat")
	}
	createEnvFile()
}

func createEnvFile() {
	path := ".env"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.WriteString(`
DB_HOST=localhost\n
DB_TYPE=postgres\n
DB_PORT=5432\n
DB_NAME=root\n
DB_PASSWORD=root\n
DB_USER=postgres\n
`)
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
	createConfigFile()
}

func createConfigFile() {
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
