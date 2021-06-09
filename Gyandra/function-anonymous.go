package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

//func blackListAdmin(name string) bool {
	//return name == "admin"
//}

//func blackListRoot(name string) bool {
	//return name == "root"
//}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("Gya", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Gya", func(name string) bool {
		return name == "root"
	})

}