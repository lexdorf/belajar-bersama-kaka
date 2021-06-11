package main

import (
    "iki/database"
    "fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
