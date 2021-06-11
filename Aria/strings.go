package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan", "Budi"))

	fmt.Println(strings.Split("Adi Pramesta", " "))

	fmt.Println(strings.ToLower("Adi Pramesta"))
	fmt.Println(strings.ToUpper("Adi Pramesta"))

	fmt.Println(strings.Trim("    Adi Pramesta   ", " "))
	fmt.Println(strings.ReplaceAll("Adi Adi Adi", "Adi", "Pramesta"))
}