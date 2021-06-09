package main

import "fmt"

func main() {
	var name = "Kurniawan"

	if name == "Gya" {
		fmt.Println("Hello Gya")
	} else if  name == "Budi" { 
		fmt.Println("Hello Budi")
	} else if  name == "Joko" { 
		fmt.Println("Hello Joko")
	}else {
		fmt.Println("Hi, Kenalan donk")
	}

	 
	if length := len(name) ;length > 5 {	
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
		
}