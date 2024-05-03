package main

import (
	"fmt"
)

type Users struct {
	name     string
	lastname string
}

func (u Users) Fullname() string {
	return u.name + " " + u.lastname
}
func main() {
	user := Users{
		name:     "Amp",
		lastname: "Sarakhanx",
	}
	Fullname := user.Fullname()
	fmt.Println(Fullname)
}
