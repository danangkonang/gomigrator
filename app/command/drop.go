package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DropRun() {
	switch os.Args[2] {
	case "--hepl":
		fmt.Println("helper")
	case "-h":
		fmt.Println("helper")
	default:
		_, err := os.Stat("migration/database/migration")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		execDrop("down", os.Args[2:])
	}
}

func execDrop(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/migration.go", typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
