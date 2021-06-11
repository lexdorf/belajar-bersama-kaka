package main

import (
	"fmt"
	"container/list"
)

func main(){
	data := list.New()

	data.PushBack("Adi")
	data.PushBack("Pramesta")
	data.PushBack("Joko")
	data.PushFront("Budi")

// dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

// dari belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}