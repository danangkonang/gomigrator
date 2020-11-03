package cmd

import (
	"fmt"
	"os"
	"strings"
)

var dirConfig = "config"

func Initial() {
	if _, err := os.Stat(dirConfig); os.IsNotExist(err) {
		os.Mkdir(dirConfig, 0700)
		fmt.Println(dirConfig + " telah dibuat")
	}
	createEnvFile()
}

func createEnvFile() {
	path := ".env"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file.WriteString(`DB_CONNECTION=postgres
DB_HOST=localhost
DB_TYPE=postgres
DB_PORT=5432
DB_NAME=testing
DB_PASSWORD=postgres
DB_USER=postgres
`)
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
	createConfigFile()
}

func createConfigFile() {
	path := dirConfig + "/config.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package config\n\n"
		writeText += "import (\n"
		writeText += `	"database/sql"`
		writeText += "\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"os"`
		writeText += "\n\n"
		writeText += `	_ "github.com/go-sql-driver/mysql"`
		writeText += "\n"
		writeText += `	"github.com/joho/godotenv"`
		writeText += "\n"
		writeText += `	_ "github.com/lib/pq"`
		writeText += "\n"
		writeText += ")\n\n"
		writeText += "func Connect() *sql.DB {\n"
		writeText += "	dbUrl := GetDbUrl()\n"
		writeText += `	connection := os.Getenv("DB_CONNECTION")`
		writeText += "\n"
		writeText += "	db, err := sql.Open(connection, dbUrl)\n"

		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "	err = db.Ping()\n"
		writeText += "	if err != nil {\n"
		writeText += "		panic(err)\n"
		writeText += "	}\n"
		writeText += "	return db\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func GetDbUrl() string {\n"
		writeText += "	e := godotenv.Load()\n"
		writeText += "	if e != nil {\n"
		writeText += "		fmt.Print(e)\n"
		writeText += "	}\n"
		writeText += `	var url = os.Getenv("DATABASE_URL")`
		writeText += "\n"
		writeText += `	if url == "" {`
		writeText += "\n"
		writeText += `		user := os.Getenv("DB_USER")`
		writeText += "\n"
		writeText += `		password := os.Getenv("DB_PASSWORD")`
		writeText += "\n"
		writeText += `		dbname := os.Getenv("DB_NAME")`
		writeText += "\n"
		writeText += `		port := os.Getenv("DB_PORT")`
		writeText += "\n"
		writeText += `		host := os.Getenv("DB_HOST")`
		writeText += "\n"
		writeText += `		localUrl := fmt.Sprintf("host=%s port=%s user=%s "+`
		writeText += "\n"
		writeText += `			"password=%s dbname=%s sslmode=disable",`
		writeText += "\n"
		writeText += `			host, port, user, password, dbname)`
		writeText += "\n"
		writeText += "		return localUrl\n"
		writeText += "	} else {\n"
		writeText += "		return url\n"
		writeText += "	}\n"

		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
	makeDirDatabase()
}

func makeDirDatabase() {
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		os.Mkdir(migrationDir, 0700)
		fmt.Println(migrationDir + " telah dibuat")
	}
	makeMaster()
}

func makeMaster() {
	tableName := "master"
	fileName := tableName + ".go"
	// fileName := "master.go"
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
		// writeMigration += "import (\n"
		// writeMigration += `	"log"`
		// writeMigration += "\n\n"
		// writeMigration += `	"github.com/danangkonang/migrasion-go-cli/config"`
		// writeMigration += "\n"
		// writeMigration += ")\n\n"
		writeMigration += "func " + strings.Title(tableName) + "(){\n"
		// writeMigration += "	db := config.Connect()\n"
		// writeMigration += "	db.Exec(`DROP TABLE " + tableName + "`)\n"
		// writeMigration += "	_, err := db.Exec(`CREATE TABLE " + tableName + "(\n"
		// writeMigration += "	id_product serial PRIMARY KEY,\n"
		// writeMigration += "	created_at TIMESTAMP NOT NULL,\n"
		// writeMigration += "	updated_at TIMESTAMP NOT NULL\n"
		// writeMigration += "	)`)\n"
		// writeMigration += "	if err != nil {\n"
		// writeMigration += "		log.Fatal(err)\n"
		// writeMigration += "	}\n"
		writeMigration += "}\n"
		file.WriteString(writeMigration)
		defer file.Close()
	}
	// fmt.Println("file berhasil dibuat", path)
}
