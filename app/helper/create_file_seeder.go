package helper

import (
	"fmt"
	"os"
	"strings"
)

type Seeder struct {
	TableName string
	Filename  string
	DirSeeder string
}

func CreateSeedFileV2(seed *Seeder) {
	// file_name := seed.Filename + "_seeder_" + GetTime() + ".go"
	file_name := GetTime() + "_seeder_" + seed.Filename + ".go"
	path := seed.DirSeeder + "/" + file_name
	// path := seed.DirSeeder + "/" + seed.Filename + "_seeder_" + GetTime() + ".go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// now := time.Now().Format("2006-01-02 15:04:05")

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
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
		writeMigration += `	"github.com/danangkonang/` + MyRootDir() + `/db/migration"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (s *Seeder) " + strings.Title(seed.Filename) + "() {\n"
		// writeMigration += "	db := config.Connect()\n"
		writeMigration += "	query := `\n"
		writeMigration += "		INSERT INTO\n"
		writeMigration += "			" + seed.TableName + " (name, created_at, updated_at)\n"
		writeMigration += "		VALUES\n"
		writeMigration += "			($1, $2, $3)\n"
		writeMigration += "	`\n"

		writeMigration += "	_, err := migration.Connection().Db.Exec(\n"
		writeMigration += "		query,\n"
		writeMigration += `		"default", time.Now(), time.Now(),`
		writeMigration += "\n"
		writeMigration += "	)\n"

		// writeMigration += "	_, err := migration.Connection().Db.Exec(`\n"
		// writeMigration += "		INSERT INTO " + seed.TableName + " (name, created_at, updated_at) VALUES\n"
		// writeMigration += "		('lorem dolor', '" + now + "','" + now + "'),\n"
		// writeMigration += "		('lorem dolor', '" + now + "','" + now + "')\n"
		// writeMigration += "	`)\n"

		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += ` 	fmt.Println("success insert table ` + file_name + `")`
		writeMigration += "	\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println("create", path)
	fmt.Println(string(green), "success")
}
