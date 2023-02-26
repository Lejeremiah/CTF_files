package main

import (
	"encoding/gob"
	"os"
)

type User struct {
	Name  string
	Path  string
	Power string
}

func encode() {
	user := User{Name: "ctfer", Path: "/tmp/a39f682dbf98f61bcbbb7b144a9bf145/", Power: "admin"}
	filePtr, _ := os.Create("user.gob")
	encoder := gob.NewEncoder(filePtr)
	encoder.Encode(&user)
}
func main() {
	encode()
}
