package main

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()

	data.PushBack("Riski")
	data.PushBack("Yana")
	data.PushBack("Ramdhani")
	data.PushFront("Boys")

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
    }
}
