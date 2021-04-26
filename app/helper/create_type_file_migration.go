package helper

import (
	"fmt"
	"os"
)

var path_type_migration = "migration/database/migration/0.core_type_migration.go"

func CreateTypeFileMigration() {
	_, err := os.Stat(path_type_migration)

	if os.IsNotExist(err) {
		var file, err = os.Create(path_type_migration)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package migration\n\n"

		writeText += "type MyMigration struct{}"
		writeText += "\n\n"

		_, err = file.WriteString(writeText)
		if isErrors(err) {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}
