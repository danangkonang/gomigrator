package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/danangkonang/gomigrator/app/helper"
	"github.com/danangkonang/gomigrator/app/model"
)

func CreateSeeder(app *model.Create) {
	if app.TableName == "" {
		fmt.Println("table name undefined")
		os.Exit(0)
	}
	if app.FileName == "" {
		app.FileName = app.TableName
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
		// app.FileName = app.FileName + fmt.Sprintf("%03d", len(newFile))
		app.FileName = app.FileName + fmt.Sprintf("%d", time.Now().Unix())
	}

	file_name := getTime() + "_seeder_" + app.FileName + ".go"
	path := DirSeeder + "/" + file_name
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		// isi file
		writeMigration := "package seeder\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"os"`
		writeMigration += "\n"
		writeMigration += `	"time"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/` + helper.MyRootDir() + `/db/migration"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (s *Seeder) " + strings.Title(app.FileName) + "() {\n"
		writeMigration += "	start := time.Now()\n"
		writeMigration += "	query := `\n"
		writeMigration += "		INSERT INTO\n"
		writeMigration += "			" + app.TableName + " (name, created_at, updated_at)\n"
		writeMigration += "		VALUES\n"
		if os.Getenv("DB_DRIVER") == "mysql" {
			writeMigration += "			(?, ?, ?)\n"
		} else {
			writeMigration += "			($1, $2, $3)\n"
		}
		writeMigration += "	`\n"
		writeMigration += "	stmt, _ := migration.Connection().Db.Prepare(query)\n"
		writeMigration += "	_, err := stmt.Exec(\n"
		writeMigration += `		"default", time.Now(), time.Now(),`
		writeMigration += "\n"
		writeMigration += "	)\n"

		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += "	duration := time.Since(start)\n"
		writeMigration += `	fmt.Println("insert table ` + file_name + `", string(migration.Green), "success", string(migration.Reset), "in", fmt.Sprintf("%.2f", duration.Seconds()), "second")`
		writeMigration += "	\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println(string(green), "success", string(white), "created", path)
}
