package main

import "fmt"

func getFullBame() (string, string, string) {
	return "Gya", "Naufal", "Pratama"
}

func main() {
	firstName, _, _ := getFullBame()
	fmt.Println(firstName)
	//fmt.Println(middleName)
	//fmt.Println(lastName)

}