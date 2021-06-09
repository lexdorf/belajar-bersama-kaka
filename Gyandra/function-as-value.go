package main

import "fmt"

func getGoodBye(nama string) string {
	return "Good Bye " + nama
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Gya")
	fmt.Println(result)
	fmt.Println(getGoodBye("Gya"))
}