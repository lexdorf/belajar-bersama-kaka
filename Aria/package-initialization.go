package main

import (
	"Aria/database"
	"fmt"
)

func main(){
	result := database.GetDatabase()
	fmt.Println(result)
}