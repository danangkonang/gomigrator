package helper

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Seeder struct {
	TableName string
	Filename  string
	DirSeeder string
}

func CreateSeedFile(seed *Seeder) {
	path := seed.DirSeeder + "/" + seed.Filename + "_seeder_" + GetTime() + ".go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	now := time.Now().Format("2006-01-02 15:04:05")

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		// isi file
		writeMigration := "package seed\n\n"
		writeMigration += "import (\n"
		writeMigration += `	"fmt"`
		writeMigration += "\n"
		writeMigration += `	"os"`
		writeMigration += "\n\n"
		writeMigration += `	"github.com/danangkonang/` + MyRootDir() + `/migration/core/config"`
		writeMigration += "\n"
		writeMigration += ")\n\n"
		writeMigration += "func (s MySeed) " + strings.Title(seed.Filename) + "() {\n"
		writeMigration += "	db := config.Connect()\n"
		writeMigration += "	_, err := db.Exec(`\n"
		writeMigration += "		INSERT INTO " + seed.TableName + " (name, created_at, updated_at) VALUES\n"
		writeMigration += "		('lorem dolor', '" + now + "','" + now + "'),\n"
		writeMigration += "		('lorem dolor', '" + now + "','" + now + "')\n"
		writeMigration += "	`)\n"
		writeMigration += "	if err != nil {\n"
		writeMigration += "		fmt.Println(err.Error())\n"
		writeMigration += "		os.Exit(0)\n"
		writeMigration += "	}\n"
		writeMigration += ` 	fmt.Println("success insert table ` + seed.Filename + `")`
		writeMigration += "	\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	fmt.Println("create", path)
	fmt.Println(string(green), "success")
}
