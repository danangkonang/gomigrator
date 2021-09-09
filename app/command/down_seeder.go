package command

import (
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/gomigrator/app/model"
)

func DownSeeder(up *model.UpDown) {
	cmd := exec.Command("go", "run", "db/bin.go", "down", "seeder", "--tables="+strings.Join(up.Tables, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
