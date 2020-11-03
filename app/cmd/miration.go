package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

var migrationDir = "database/migration"

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
	path := migrationDir + "/" + fileName
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		// isi file
		// file.WriteString("package database\n func " + tableName + "(){\n}")
		writeMigration := "package database\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"log"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/migrasion-go-cli/config"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func " + strings.Title(tableName) + "(){\n"
		writeMigration += "	db := config.Connect()\n"
		writeMigration += "	db.Exec(`DROP TABLE " + tableName + "`)\n"
		writeMigration += "	_, err := db.Exec(`CREATE TABLE " + tableName + "(\n"
		writeMigration += "	id_product serial PRIMARY KEY,\n"
		writeMigration += "	created_at TIMESTAMP NOT NULL,\n"
		writeMigration += "	updated_at TIMESTAMP NOT NULL\n"
		writeMigration += "	)`)\n"
		writeMigration += "	if err != nil {\n"
		writeMigration += "		log.Fatal(err)\n"
		writeMigration += "	}\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
}
