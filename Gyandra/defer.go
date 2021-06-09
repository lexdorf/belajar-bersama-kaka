package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAppLication(value int) {
	defer logging()
	fmt.Println("Run appLication")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runAppLication(10)

}