package main

import (
	"Aria/helper"
	"fmt"
)

func main(){
	helper.SayHello("Adi")
	// helper.sayGoodbye("Adi") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}