package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. "+ man.Name

}

func main() {
	Gya := Man{"Gya"}
	Gya.Married()

	fmt.Println(Gya.Name)

}