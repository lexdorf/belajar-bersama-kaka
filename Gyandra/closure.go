package main

import "fmt"

func main() {
	name := "Gya"
	counter := 0

	increment := func() {
		name := "Nofal"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}