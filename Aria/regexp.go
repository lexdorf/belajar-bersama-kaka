package main

import (
	"fmt"
	"regexp"
)

func main(){
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("abi"))
	fmt.Println(regex.MatchString("ari"))
	fmt.Println(regex.MatchString("aLi"))

	search := regex.FindAllString("adi, ali, ari, aBi", -1)
	fmt.Println(search)
}