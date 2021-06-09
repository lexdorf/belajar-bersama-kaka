package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main() {
	var Address1 Address = Address{"Bandung", "Jawa Tengah", "Indonesia"}
	//var Address4 *Address = &Address{"Subang","Jawa Barat", "Indonesia"}
	var Address2 *Address = &Address1
	var Address3 *Address = &Address1

	Address2.City = "Cicalengka"

	*Address2 = Address{"Subang", "Jawa Barat", "Indonesia"}

	fmt.Println(Address1)
	fmt.Println(Address2)
	fmt.Println(Address3)

	var Address4 *Address = new(Address)
	Address4.City = "Sumatra"
	fmt.Println(Address4)

	var alamat = Address {
		City: "Bandung",
		Province: "Jawa Barat",
		Country: "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}