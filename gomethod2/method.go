package main

import (
	"fmt"
)

type Canspeake interface {
	Sound() string
}
type Person struct {
	Name string
}
type Dog struct {
	Name string
}

func (p Person) Sound() string {
	return "Kuay"
}

func (d Dog) Sound() string {
	return "Huff"
}

func SoundOf(c Canspeake, name string) {
	fmt.Println(name + " Sound like " + c.Sound())
}

func main() {
	person := Person{Name: "Amp"}
	SoundOf(person, person.Name)
	dog := Dog{Name: "Bo"}
	SoundOf(dog, dog.Name)
}
