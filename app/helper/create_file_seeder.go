package helper

import (
	"fmt"
	"os"
	"strings"
)

func CreateSeedFile(rootDir, seederDir string) {
	tableName := os.Args[3]
	fileName := tableName + "_" + GetTime() + ".go"
	path := seederDir + "/" + fileName
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
		writeMigration := "package seed\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"log"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/` + rootDir + `/migration/app/config"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func " + strings.Title(tableName) + "(){\n"
		writeMigration += "	db := config.Connect()\n"
		// writeMigration += "	db.Exec(`DROP TABLE " + tableName + "`)\n"
		writeMigration += "	_, err := db.Exec(`INSERT INTO " + tableName + " (created_at,updated_at)VALUES\n"
		// writeMigration += "	id_product serial PRIMARY KEY,\n"
		// writeMigration += "	created_at TIMESTAMP NOT NULL,\n"
		writeMigration += "	('2006-01-02 15:04:05','2006-01-02 15:04:05'),\n"
		writeMigration += "	('2006-01-02 15:04:05','2006-01-02 15:04:05')\n"
		writeMigration += "	`)\n"
		writeMigration += "	if err != nil {\n"
		writeMigration += "		log.Fatal(err)\n"
		writeMigration += "	}\n"
		writeMigration += `fmt.Println("success insert table ` + fileName + `")`
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println("create", path)
}
