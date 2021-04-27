package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DropTableRun() {
	if len(os.Args[2:]) > 0 {
		// fmt.Println("manual")
		if os.Args[2] == "--help" {
			fmt.Println("helper")
			os.Exit(0)
		}
		if os.Args[2] == "-h" {
			fmt.Println("helper")
			os.Exit(0)
		}
		// fmt.Println("helper lain")
		_, err := os.Stat("migration/database/migration")
		if err != nil {
			fmt.Println("no initial direktori")
			os.Exit(0)
		}
		execDrop("drop", os.Args[2:])
	}
	if len(os.Args[2:]) == 0 {
		// fmt.Println("auto")
		execDrop("drop", os.Args[2:])
	}
}

func execDrop(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/main.go", typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
