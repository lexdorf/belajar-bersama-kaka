package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer {
		Name: "Joko",
		Address: "Bandung",
		Age: 35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
	

}