package helper

import (
	"fmt"
	"os"
	"strings"
)

var (
	red   = "\033[31m"
	green = "\033[32m"
	blue  = "\033[34m"
)

func CreateMigrationFile(name, migrationDir string) {
	tableName := os.Args[3]
	fileName := tableName + "_" + GetTime() + ".go"
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
		writeMigration := "package migration\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"log"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/` + name + `/migration/app/config"`
		// writeMigration += `	"github.com/danangkonang/migrasion-go-cli/config"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func " + strings.Title(tableName) + "(){\n"
		writeMigration += "	db := config.Connect()\n"
		writeMigration += "	db.Exec(`DROP TABLE " + tableName + "`)\n"
		writeMigration += "	_, err := db.Exec(`CREATE TABLE " + tableName + "(\n"
		writeMigration += "	id serial PRIMARY KEY,\n"
		writeMigration += "	created_at TIMESTAMP NOT NULL,\n"
		writeMigration += "	updated_at TIMESTAMP NOT NULL\n"
		writeMigration += "	)`)\n"
		writeMigration += "	if err != nil {\n"
		writeMigration += "		log.Fatal(err)\n"
		writeMigration += "	}\n"
		writeMigration += `fmt.Println("success create table ` + fileName + `")`
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println("create", path)
	fmt.Println(string(green), "success")
}
