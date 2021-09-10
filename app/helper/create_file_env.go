package helper

import (
	"fmt"
	"os"

	"github.com/danangkonang/gomigrator/app/model"
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
			"DB_DRIVER=%s\nDB_HOST=%s\nDB_PORT=%d\nDB_NAME=%s\nDB_USER=%s\nDB_PASSWORD=%s\n",
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
