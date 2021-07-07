package helper

import (
	"fmt"
	"os"
)

var path_type_execution = "migration/core/execusion/type_execution.go"

func CreateTypeFileExecution() {
	_, err := os.Stat(path_type_execution)

	if os.IsNotExist(err) {
		var file, err = os.Create(path_type_execution)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		writeText := "package execusion\n\n"

		writeText += "type Tables struct {\n"
		writeText += "	NameTable []string\n"
		writeText += "}\n\n"

		// writeText += "type MyTp struct{}"
		// writeText += "\n\n"

		_, err = file.WriteString(writeText)
		if isErrors(err) {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		defer file.Close()
	}

}
