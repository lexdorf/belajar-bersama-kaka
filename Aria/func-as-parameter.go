package main

import "fmt"

type Filter func(string) string

func sayeHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Bangsat" {
		return "..."
	} else {
		return name
	}
}

func main(){
	sayeHelloWithFilter("Adi", spamFilter)
	sayeHelloWithFilter("Bangsat", spamFilter)
}