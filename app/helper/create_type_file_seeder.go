package helper

import (
	"fmt"
	"os"
)

var path_type_seed = "migration/database/seed/0.core_type_seed.go"

func CreateTypeFileSeeder() {
	_, err := os.Stat(path_type_seed)

	if os.IsNotExist(err) {
		var file, err = os.Create(path_type_seed)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package seed\n\n"

		writeText += "type MySeed struct{}"
		writeText += "\n\n"

		_, err = file.WriteString(writeText)
		if isErrors(err) {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}
}
