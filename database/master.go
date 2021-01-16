package database

import (
	"github.com/danangkonang/migrasion-go-cli/database/migration"
	"github.com/danangkonang/migrasion-go-cli/database/seed"
)

func Master() {
	migration.Users()
	migration.Products()
	migration.Sub_categories()
	migration.Categories()
	migration.Sub_items()
	migration.User_profiles()
	migration.Kelurahan()
	migration.Kecamatan()
	migration.Kabupaten()
	migration.Provinsi()
	migration.Messages()
	migration.Notif()
	migration.News()
}

func Seeds() {
	seed.Categories()
	seed.Sub_categories()
	seed.Sub_items()
	seed.Provinsi()
	seed.Kabupaten()
	seed.Kecamatan1to3()
	seed.Kecamatan4to6()
	seed.Kecamatan7to9()
	seed.Kecamatan10to12()
	seed.Kecamatan13to15()
	seed.Kecamatan16to18()
	seed.Kecamatan19to21()
	seed.Kecamatan22to25()
	seed.Kecamatan26to28()
	seed.Kecamatan29to34()
}
