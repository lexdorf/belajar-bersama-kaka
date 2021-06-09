package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

//func (man Man) Married() {
//	man.Name = "Mr. " + man.Name
//	fmt.Println("Di Method", man.Name)
//}

func main(){
	pramesta := Man{"Pramesta"}
	pramesta.Married()

	fmt.Println(pramesta.Name)
}