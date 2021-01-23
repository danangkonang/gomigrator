package helper

import (
	"fmt"
	"os"
)

func CreateMainFile(thisDir, dirMigration string) {
	path := dirMigration + "/migration.go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)
	// buat file baru jika belum ada

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		writeText := "package main\n\n"
		writeText += "import (\n"
		// writeText += `	"os"`
		// writeText += "\n\n"
		// writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/command"`
		// "github.com/danangkonang/database-migration/app/helper"
		// writeText += "\n"
		writeText += `	//"github.com/danangkonang/` + thisDir + `/migration/app/helper"`
		writeText += "\n"
		writeText += `	"github.com/danangkonang/` + thisDir + `/migration/app/execusion"`
		writeText += "\n"
		writeText += ")\n\n"

		// function
		writeText += "func main() {\n"
		// writeText += "	arrCmd := os.Args[1:]\n"

		// writeText += "	if len(arrCmd) == 0 {\n"
		// writeText += "		helper.PrintHelper()\n"
		// writeText += "		return\n"
		// writeText += "	}\n"

		writeText += "	runCmd()\n"
		writeText += "}\n"
		writeText += "\n"

		writeText += "func runCmd() {\n"
		// writeText += "	usrCmd := os.Args[1]\n"
		// writeText += "	switch usrCmd {\n"

		// writeText += `	case "-h":`
		// writeText += "\n"
		writeText += "	//helper.PrintHelper() \n"
		writeText += "	execusion.RuningMigration()\n"
		// writeText += `	case "--help":`
		// writeText += "\n"
		// writeText += "		helper.PrintHelper() \n"
		// writeText += "		break\n"
		// writeText += `	case "-v":`
		// writeText += "\n"
		// writeText += "			helper.PrintVersion() \n"
		// writeText += "		break\n"
		// writeText += `	case "--version":`
		// writeText += "\n"
		// writeText += "		helper.PrintVersion() \n"
		// writeText += "		break\n"
		// writeText += `	case "init":`
		// writeText += "\n"
		// writeText += " 		cmd.Initial()\n" //create .env config database
		// writeText += "		break\n"
		// writeText += `	case "danang":`
		// writeText += "\n"
		// writeText += "		cmd.TypeMigration()\n"
		// writeText += "		break\n"
		// writeText += `	case "create":`
		// writeText += "\n"
		// writeText += "		cmd.TypeMigration()\n" //go run main.go create migration [name file]
		// writeText += "		break\n"
		// writeText += `	case "run":`
		// writeText += "\n"
		// writeText += "		cmd.TipeRun()\n" //go run main.go run migration [name file]
		// writeText += "		break\n"

		// writeText += `	case "back":`
		// writeText += "\n"
		// writeText += "		cmd.TypeBack()\n"
		// writeText += "		break\n"
		// writeText += "	default:\n"
		// writeText += "		helper.PrintHelper()\n"
		// writeText += "	}\n"
		writeText += "}\n"

		file.WriteString(writeText)
		defer file.Close()
	}
}
