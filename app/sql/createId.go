package sql

import (
	"fmt"
	"strconv"
)

type Database struct {
	Name   string
	Length string
}

func Unix(Nama string, Panjang int) string {
	// {
	// 	unix: func uuid()  string{
	// 		fmt.Println("unix")
	// 		return "kjkjk"
	// 	}
	// },
	// db.Name = "id"
	pjg := strconv.Itoa(Panjang)
	fmt.Println(pjg)
	// id serial PRIMARY KEY
	return Nama + " serial PRIMARY KEY,"
}

func String(Nama string, Panjang int) string {
	// category VARCHAR (225) NULL
	pjg := strconv.Itoa(Panjang)
	return Nama + " VARCHAR (" + pjg + ") NULL,"
}

func TimeStame(Nama string) string {
	// created_at TIMESTAMP NOT NULL
	return Nama + " TIMESTAMP NOT NULL,"
}
