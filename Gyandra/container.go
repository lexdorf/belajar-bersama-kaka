package main

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()

	data.PushBack("Gya")
	data.PushBack("Naufal")
	data.PushBack("Pratama")
	data.PushFront("Joko")

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}