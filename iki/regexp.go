package main

import (
	"fmt"
	"regexp"
)

func main(){
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("aiski"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("aini"))
	fmt.Println(regex.MatchString("aiki"))

	search := regex.FindAllString("aiski, aiki, aini, adi", -1)
	fmt.Println(search)
}




