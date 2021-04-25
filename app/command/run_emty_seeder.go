package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func EmtySeederData() {
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
		execEmtySeederData("down", os.Args[2:])
	}
}

func execEmtySeederData(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/main.go", typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
