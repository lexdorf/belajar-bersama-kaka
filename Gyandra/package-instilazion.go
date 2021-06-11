package main

import(
	"Gyandra/database"
"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

}