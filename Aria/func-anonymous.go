package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, balckList Blacklist) {
	if balckList(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	balckList := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", balckList)
	registerUser("Adi", balckList)

	registerUser("bro", func(name string) bool {
		return name == "bro"
	})
	registerUser("Adi", func(name string) bool {
		return name == "bro"
	})
}