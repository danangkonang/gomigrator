package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DropTableRun() {
	// fmt.Println(os.Args[1:])
	switch os.Args[1] {
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
		execDrop("drop", os.Args[2:]) // delete table
	}
}

func execDrop(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/main.go", typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
