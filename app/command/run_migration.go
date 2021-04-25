package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/migrasion-go-cli/app/helper"
)

func MigrationsRun() {
	// fullCommand := os.Args[1:]
	// if len(fullCommand) == 1 {
	// 	helper.PrintHelper()
	// 	return
	// }
	// fmt.Println(os.Args[2])
	switch os.Args[2] {
	case "--hepl":
		fmt.Println("helper")
	case "-h":
		fmt.Println("helper")
	case "migration":
		// fmt.Println(len(os.Args[3:]))
		execMigration("run", "migration", os.Args[3:])
	case "-m":
		execMigration("run", "migration", os.Args[3:])
	case "seeder":
		// fmt.Println(len(os.Args[3:]))
		execMigration("run", "seeder", os.Args[3:])
	case "-s":
		execMigration("run", "seeder", os.Args[3:])
	default:
		helper.ErrorCommand(os.Args[2])
	}
}

func execMigration(firstFlag, typeFlag string, second []string) {
	// fmt.Println(len(strings.Join(second, ",")))
	cmd := exec.Command("go", "run", "migration/main.go", firstFlag, typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// func runExecusin(typeFlag string) {
// 	cmd := exec.Command("go", "run", "migration/migration.go", typeFlag)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }
