package main

import (
    "flag"
    "fmt"
)

func main() {
	var host *string = flag.String("host", "Localhost", "Put your data host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")

    flag.Parse()

    fmt.Println("Host : ", *host)
    fmt.Println("User : ", *user)
    fmt.Println("Password : ", *password)
}
