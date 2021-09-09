package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/gomigrator/app/model"
)

var (
	red   = "\033[31m"
	green = "\033[32m"
	blue  = "\033[34m"
	white = "\033[0;37m"
)

func getTime() string {
	t := time.Now()
	waktu := t.Format("2006_01_02_150405")
	return waktu
}

func CreateMigtaion(app *model.Create) {
	if app.TableName == "" {
		fmt.Println("table name null")
		os.Exit(0)
	}
	files, erro := ioutil.ReadDir(DirSeeder)
	if erro != nil {
		fmt.Println(erro.Error())
		os.Exit(0)
	}
	newFile := []string{}
	for _, file := range files {
		filename := file.Name()
		list := strings.Split(filename, "_seeder_")
		if list[0] != "0.go" {
			name := list[1]
			tb_name := strings.Split(name, ".go")
			if app.TableName == tb_name[0] {
				newFile = append(newFile, tb_name[0])
			}
		}
	}
	if len(newFile) > 0 {
		fmt.Println("duplicate table name")
		os.Exit(0)
	}
	file_name := getTime() + "_migration_" + app.TableName + ".go"
	path := DirMigration + "/" + file_name
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
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (m *Migration) " + strings.Title(app.TableName) + "() {\n"

		writeMigration += "	query := `\n"
		writeMigration += "		CREATE TABLE " + app.TableName + "(\n"
		writeMigration += "			id serial PRIMARY KEY,\n"
		writeMigration += "			name VARCHAR (225) NOT NULL,\n"
		writeMigration += "			created_at TIMESTAMP NOT NULL,\n"
		writeMigration += "			updated_at TIMESTAMP NOT NULL\n"
		writeMigration += "		)\n"
		writeMigration += "	`\n"

		writeMigration += "	_, err := Connection().Db.Exec(query)\n"

		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += `  fmt.Println(string(Green), "success", string(Reset), "create table ` + file_name + `")`
		writeMigration += "\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println(string(green), "success", string(white), "created", path)
}
