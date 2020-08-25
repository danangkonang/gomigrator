// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// var dir = "database"
// var path = dir + "/" + getTime() + "-test.go"

// func isError(err error) bool {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return (err != nil)
// }

// func getTime() string {
// 	t := time.Now()
// 	waktu := t.Format("2006-01-02-150405")
// 	return waktu
// }

// func createFile() {
// 	// deteksi apakah file sudah ada
// 	_, err := os.Stat(path)

// 	// buat file baru jika belum ada
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		file.WriteString("package database")
// 		defer file.Close()
// 	}
// 	fmt.Println("file berhasil dibuat", path)
// }

// func createFolder() {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		os.Mkdir(dir, 0700)
// 		fmt.Println(dir + " telah dibuat")
// 	}
// }

// func main() {
// 	createFolder()
// 	createFile()
// 	fmt.Println(getTime())
// }

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var dir = "database"

// var path = dir + "/" + getTime() + "-test.go"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func cli() {
	name := flag.String("migrate", "", "make file migration 'konang -migrate users'")
	// seed := flag.String("migrate", "", "make file migration 'konang -migrate users'")
	flag.Parse()
	if len(*name) == 0 {
		os.Exit(1)
	}
	createFolder(name)
}

func createFolder(name *string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
		fmt.Println(dir + " telah dibuat")
	}
	createFile(name)
}

func createFile(name *string) {
	path := dir + "/" + getTime() + "-" + *name + ".go"
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		file.WriteString("package database\n")
		defer file.Close()
	}
	fmt.Println("file berhasil dibuat", path)
}

func getTime() string {
	t := time.Now()
	waktu := t.Format("2006-01-02-150405")
	return waktu
}

func main() {
	cli()
}
