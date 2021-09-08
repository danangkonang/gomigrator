package command

import (
	"os"
	"os/exec"
	"strings"

	"github.com/danangkonang/migration-go-cli/app/model"
)

func DownSeeder(up *model.UpDown) {
	cmd := exec.Command("go", "run", "db/bin.go", "down", "seeder", "--tables="+strings.Join(up.Tables, ","))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
