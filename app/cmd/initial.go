package cmd

import (
	"fmt"
	"os"
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
		file.WriteString(`
DB_CONNECTION=mysql
DB_HOST=localhost
DB_TYPE=postgres
DB_PORT=5432
DB_NAME=root
DB_PASSWORD=root
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
		writeText += `		user := os.Getenv("DBUSER")`
		writeText += "\n"
		writeText += `		password := os.Getenv("DBPASSWORD")`
		writeText += "\n"
		writeText += `		dbname := os.Getenv("DBNAME")`
		writeText += "\n"
		writeText += `		port := os.Getenv("DBPORT")`
		writeText += "\n"
		writeText += `		host := os.Getenv("DBHOST")`
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
}
