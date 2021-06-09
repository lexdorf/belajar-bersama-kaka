package main

import "fmt"

func endApp(){
	massage := recover()
	if massage != nil {
		fmt.Println("Error dengan massage:", massage)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main(){
	runApp(true)
}