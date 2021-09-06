package helper

import (
	"fmt"
	"os"

	"github.com/danangkonang/migration-go-cli/app/model"
)

func NewCreateEnvFile(app *model.Init) {
	path := ".env"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		file_env := fmt.Sprintf(
			"DB_DRIVER=%s\nDB_HOST=%s\nDB_PORT=%d\nDB_NAME=%s\nDB_PASSWORD=%s\nDB_USER=%s",
			app.Driver,
			app.Host,
			app.Port,
			app.DAtabase,
			app.User,
			app.Password,
		)
		file.WriteString(file_env)
		defer file.Close()
	}

}

func CreateEnvFile(driver string) {
	path := ".env"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		var port int
		if driver == "mysql" {
			port = 3306
		} else {
			port = 5432
		}
		file_env := fmt.Sprintf(
			"DB_CONNECTION=%s\nDB_HOST=localhost\nDB_TYPE=%s\nDB_PORT=%d\nDB_NAME=default\nDB_PASSWORD=%s\nDB_USER=%s",
			driver,
			driver,
			port,
			driver,
			driver,
		)
		fmt.Println(file_env)
		// 		file.WriteString(`DB_CONNECTION=postgres
		// DB_HOST=localhost
		// DB_TYPE=postgres
		// DB_PORT=5432
		// DB_NAME=default
		// DB_PASSWORD=postgres
		// DB_USER=postgres
		// `)
		file.WriteString(file_env)
		defer file.Close()
	}

}
