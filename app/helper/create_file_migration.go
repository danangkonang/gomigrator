package helper

import (
	"fmt"
	"os"
	"strings"
)

type Migration struct {
	TableMigration string
	DirMigration   string
}

var (
	// red   = "\033[31m"
	green = "\033[32m"
	// blue  = "\033[34m"
)

func CreateMigrationFile(migration *Migration) {
	table_name := migration.TableMigration
	file_name := table_name + "_migration_" + GetTime() + ".go"
	path := migration.DirMigration + "/" + file_name
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		// isi file
		writeMigration := "package migration\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"os"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/` + MyRootDir() + `/migration/app/config"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (m MyMigration) " + strings.Title(table_name) + "() {\n"
		writeMigration += "	db := config.Connect()\n"
		writeMigration += "	_, err := db.Exec(`\n"
		writeMigration += "		CREATE TABLE " + table_name + "(\n"
		writeMigration += "			id serial PRIMARY KEY,\n"
		writeMigration += "			name VARCHAR (225) NOT NULL,\n"
		writeMigration += "			created_at TIMESTAMP NOT NULL,\n"
		writeMigration += "			updated_at TIMESTAMP NOT NULL\n"
		writeMigration += "		);\n"
		writeMigration += "	`)\n"
		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += `  fmt.Println("success create table ` + file_name + `")`
		writeMigration += "\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println("create", path)
	fmt.Println(string(green), "success")
}
