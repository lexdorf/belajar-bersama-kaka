package main

import "fmt"
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	person := NewMap("Adi")
	fmt.Println(person)

//	var person map[string]string = NewMap("")

//	if person == nil {
//		fmt.Println("Data Kosong")
//	} else {
//		fmt.Println(person)
//	}
}