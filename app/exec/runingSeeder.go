package exec

import "github.com/danangkonang/migrasion-go-cli/database"

func RunSeeder() {
	database.Seeds()
}
