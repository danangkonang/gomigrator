package helper

import (
	"fmt"
	"os"

	"github.com/danangkonang/gomigrator/app/model"
)

var (
	zero_migration = "db/migration/0.go"
)

func NewCreateZeroMigration(app *model.Init) {
	_, err := os.Stat(zero_migration)

	if os.IsNotExist(err) {
		var file, err = os.Create(zero_migration)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package migration\n\n"

		writeText += "import (\n"
		writeText += `	"database/sql"`
		writeText += "\n"
		writeText += `	"fmt"`
		writeText += "\n"
		// writeText += `	"log"`
		// writeText += "\n"
		writeText += `	"os"`
		writeText += "\n"
		writeText += `	"time"`
		writeText += "\n\n"
		switch app.Driver {
		case "mysql":
			writeText += `	_ "github.com/go-sql-driver/mysql"`
		case "postgres":
			writeText += `	_ "github.com/lib/pq"`
		default:
			writeText += `	_ "github.com/lib/pq"`
		}
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type Migration struct{}\n\n"

		writeText += "type DB struct {\n"
		writeText += "	Db *sql.DB\n"
		writeText += "}\n\n"

		writeText += "var (\n"
		writeText += `	Green = "\033[32m"`
		writeText += "\n"
		writeText += `	Reset = "\033[0m"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "func Connection() *DB {\n"
		writeText += "	var connection string\n"

		writeText += "	psql := fmt.Sprintf(\n"
		writeText += `		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",`
		writeText += "\n"
		writeText += `		os.Getenv("DB_HOST"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_PORT"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_USER"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_PASSWORD"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_NAME"),`
		writeText += "\n"
		writeText += "	)\n"

		writeText += "	mysql := fmt.Sprintf(\n"
		writeText += `		"%s:%s@tcp(%s:%s)/%s",`
		writeText += "\n"
		writeText += `		os.Getenv("DB_USER"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_PASSWORD"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_HOST"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_PORT"),`
		writeText += "\n"
		writeText += `		os.Getenv("DB_NAME"),`
		writeText += "\n"
		writeText += "	)\n"

		writeText += `	switch os.Getenv("DB_DRIVER") {`
		writeText += "\n"
		writeText += `	case "postgres":`
		writeText += "\n"
		writeText += "		connection = psql\n"
		writeText += `	case "mysql":`
		writeText += "\n"
		writeText += "		connection = mysql\n"
		writeText += "	default:\n"
		writeText += "		connection = psql\n"
		writeText += "	}\n"

		writeText += `	db, err := sql.Open(os.Getenv("DB_DRIVER"), connection)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		fmt.Println(err.Error())\n"
		writeText += "		os.Exit(0)\n"
		writeText += "	}\n"
		writeText += "	db.SetMaxOpenConns(100)\n"
		writeText += "	db.SetMaxIdleConns(100)\n"
		writeText += "	db.SetConnMaxLifetime(5 * time.Minute)\n"
		writeText += "	return &DB{Db: db}\n"
		writeText += "}\n"

		_, err = file.WriteString(writeText)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}

func CreateZeroMigration() {
	_, err := os.Stat(zero_migration)

	if os.IsNotExist(err) {
		var file, err = os.Create(zero_migration)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package migration\n\n"

		writeText += "import (\n"
		writeText += `	"database/sql"`
		writeText += "\n"
		writeText += `	"fmt"`
		writeText += "\n"
		writeText += `	"log"`
		writeText += "\n"
		writeText += `	"time"`
		writeText += "\n\n"
		writeText += `	_ "github.com/lib/pq"`
		writeText += "\n"
		writeText += ")\n\n"

		writeText += "type Migration struct{}\n\n"

		writeText += "type DB struct {\n"
		writeText += "	Db *sql.DB\n"
		writeText += "}\n\n"

		writeText += "func Connection() *DB {\n"
		writeText += "	connection := fmt.Sprintf(\n"
		writeText += `		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",`
		writeText += "\n"
		writeText += `		"localhost",`
		writeText += "\n"
		writeText += "		5432,\n"
		writeText += `		"postgres",`
		writeText += "\n"
		writeText += `		"postgres",`
		writeText += "\n"
		writeText += `		"default",`
		writeText += "\n"
		writeText += "	)\n"
		writeText += `	db, err := sql.Open("postgres", connection)`
		writeText += "\n"
		writeText += "	if err != nil {\n"
		writeText += "		log.Panic(err.Error())\n"
		writeText += "	}\n"
		writeText += "	db.SetMaxOpenConns(100)\n"
		writeText += "	db.SetMaxIdleConns(100)\n"
		writeText += "	db.SetConnMaxLifetime(5 * time.Minute)\n"
		writeText += "	return &DB{Db: db}\n"
		writeText += "}\n"

		_, err = file.WriteString(writeText)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}
