package seed

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Provinsi() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO provinsi (provinsi)VALUES
		('BANTEN'),
		('DKI JAKARTA'),
		('JAWA BARAT'),
		('JAWA TENGAH'),
		('DI YOGYAKARTA'),
		('JAWA TIMUR'),
		('BALI'),
		('NANGGROE ACEH DARUSSALAM (NAD)'),
		('SUMATERA UTARA'),
		('SUMATERA BARAT'),
		('RIAU'),
		('KEPULAUAN RIAU'),
		('JAMBI'),
		('BENGKULU'),
		('SUMATERA SELATAN'),
		('BANGKA BELITUNG'),
		('LAMPUNG'),
		('KALIMANTAN BARAT'),
		('KALIMANTAN TENGAH'),
		('KALIMANTAN SELATAN'),
		('KALIMANTAN TIMUR'),
		('KALIMANTAN UTARA'),
		('SULAWESI BARAT'),
		('SULAWESI SELATAN'),
		('SULAWESI TENGGARA'),
		('SULAWESI TENGAH'),
		('GORONTALO'),
		('SULAWESI UTARA'),
		('MALUKU'),
		('MALUKU UTARA'),
		('NUSA TENGGARA BARAT (NTB)'),
		('NUSA TENGGARA TIMUR (NTT)'),
		('PAPUA BARAT'),
		('PAPUA')
	`)
	if err != nil {
		log.Fatal(err)
	}
}
