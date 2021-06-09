package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHai(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main(){
	var adi Customer
	adi.Name = "Adi"
	adi.Address = "Indonesia"
	adi.Age = 16

	adi.sayHai("Joko")


//	fmt.Println(adi)
//	fmt.Println(adi.Name)
//	fmt.Println(adi.Address)
//	fmt.Println(adi.Age)

//	joko := Customer{
//		Name: "Joko",
//		Address: "Bandung",
//		Age: 35,
//	}
//	fmt.Println(joko)
//
//	budi := Customer{"Budi", "Jakarta", 45}
//	fmt.Println(budi)
}