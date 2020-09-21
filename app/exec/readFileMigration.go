package exec

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadeFileInFolder() {
	files, err := ioutil.ReadDir("database")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
