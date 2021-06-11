package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Gya Naufal", "Gya"))
	fmt.Println(strings.Contains("Gya Naufal", "Budi"))

	fmt.Println(strings.Split("Gya Naufal", " "))

	fmt.Println(strings.ToLower("Gya Naufal"))
	fmt.Println(strings.ToUpper("Gya Naufal"))
	fmt.Println(strings.ToTitle("Gya Naufal"))

	fmt.Println(strings.Trim("      Gya Naufal     ", " "))
	fmt.Println(strings.ReplaceAll("Gya Gya Gya Gya", "Gya", "Budi"))

}