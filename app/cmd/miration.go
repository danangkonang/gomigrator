package cmd

import (
	"fmt"
	"os"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var migrationDir = "database"

func CreateMigration() {
	fullCmd := os.Args[1:]
	if len(fullCmd) < 3 {
		helper.PrintHelper()
		return
	}
	createFolderDatabase() //var migrationDir = "database"
}

func createFolderDatabase() {
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		os.Mkdir(migrationDir, 0700)
		fmt.Println(migrationDir + " telah dibuat")
	}
	createFilemigration()
}

func createFilemigration() {
	tableName := os.Args[3]
	fileName := helper.GetTime() + "_" + tableName + ".go"
	path := dir + "/" + fileName
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		// isi file
		file.WriteString("package database\n")
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
}
