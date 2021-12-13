package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danangkonang/gomigrator/app/model"
	"github.com/danangkonang/gomigrator/templates"
)

var (
	// red   = "\033[31m"
	Green = "\033[32m"
	// blue  = "\033[34m"
	white = "\033[0;37m"
)

// func getTime() string {
// 	t := time.Now()
// 	waktu := t.Format("2006_01_02_150405")
// 	return waktu
// }

func CreateUnixNumber(num int) string {
	new_number := fmt.Sprintf("%04d", num)
	return new_number
}

func CreateMigtaion(app *model.Create) {
	if app.TableName == "" {
		fmt.Println("unknow migration table name or run \n\"gomigrator create migration --help\"")
		os.Exit(0)
	}
	files, erro := ioutil.ReadDir(DirMigration)
	if erro != nil {
		hlp := &templates.Helper{
			Error: "unknow folder migration",
		}
		templates.PrintTemplate(templates.ErrorTmp, hlp)
		os.Exit(0)
	}
	newFile := []string{}
	for _, file := range files {
		filename := file.Name()
		if filename != "0.go" {
			list := strings.Split(filename, "_migration_")
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
	file_name := CreateUnixNumber(len(files)) + "_migration_" + app.TableName + ".go"
	path := DirMigration + "/" + file_name
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		// isi file
		writeMigration := "package migration\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"os"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (m *Migration) Up" + strings.Title(app.TableName) + "() {\n"
		writeMigration += "	query := `\n"
		writeMigration += "		CREATE TABLE " + app.TableName + "(\n"
		if os.Getenv("DB_DRIVER") == "mysql" {
			writeMigration += "			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,\n"
			writeMigration += "			name VARCHAR (225) NOT NULL,\n"
			writeMigration += "			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,\n"
			writeMigration += "			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP\n"
		} else {
			writeMigration += "			id serial PRIMARY KEY,\n"
			writeMigration += "			name VARCHAR (225) NOT NULL,\n"
			writeMigration += "			created_at TIMESTAMP NOT NULL,\n"
			writeMigration += "			updated_at TIMESTAMP NOT NULL\n"
		}
		writeMigration += "		)\n"
		writeMigration += "	`\n"

		writeMigration += "	_, err := Connection().Db.Exec(query)\n"

		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += `  fmt.Println(string(Green), "success", string(Reset), "up ` + file_name + `")`
		writeMigration += "\n"
		writeMigration += "}\n\n"

		writeMigration += "func (m *Migration) Down" + strings.Title(app.TableName) + "() {\n"
		writeMigration += "	query := `DROP TABLE " + app.TableName + "`\n"

		writeMigration += "	_, err := Connection().Db.Exec(query)\n"

		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += `  fmt.Println(string(Green), "success", string(Reset), "down ` + file_name + `")`
		writeMigration += "\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println(string(Green), "success", string(white), "created", path)
}
