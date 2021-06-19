package main

import (
	"fmt"
	"regexp"
)

func main(){
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("aiki"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("ani"))
	fmt.Println(regex.MatchString("ayi"))

	search := regex.FindAllString("aiki, adi, ani, ayi", -1)
	fmt.Println(search)
}




