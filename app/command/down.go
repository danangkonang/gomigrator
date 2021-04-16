package command

import (
	"os"
	"os/exec"
	"strings"
)

func Down(typeFlag string, second []string) {
	cmd := exec.Command("go", "run", "migration/migration.go", typeFlag, strings.Join(second, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
