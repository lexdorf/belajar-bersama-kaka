package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	 fmt.Println("Hello", firstName, lastName)
}

func main () {
	firstName := "Gya"
	sayHelloTo(firstName, "Naufal")
	sayHelloTo("Budi", "Kurniawan")


}