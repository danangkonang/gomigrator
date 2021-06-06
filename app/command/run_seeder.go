package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/helper"
)

func SeederRun() {
	// fullCommand := os.Args[1:]
	// if len(fullCommand) == 1 {
	// 	helper.PrintHelper()
	// 	return
	// }
	subCommand := os.Args[2]
	switch subCommand {
	case "--hepl":
		fmt.Println("helper")
	case "-h":
		fmt.Println("helper")
	// case "migration":
	// 	// runExecusin("migration")
	// 	runExecusin("migration", os.Args[3:])
	// case "-m":
	// 	// runExecusin("migration")
	// 	runExecusin("migration", os.Args[3:])
	case "seeder":
		// runExecusin("seeder")
		runExecusin("seeder", os.Args[2:])
	case "-s":
		// runExecusin("seeder")
		runExecusin("seeder", os.Args[2:])
	default:
		helper.ErrorCommand(os.Args[2])
	}
}

func runExecusin(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/main.go", typeFlag, strings.Join(second, ","))
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
