package main

import(
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("gya"))
	fmt.Println(regex.MatchString("elang"))
	fmt.Println(regex.MatchString("iDoer"))

	search := regex.FindAllString("gya yang edo eto eyo eki", 2)
	fmt.Println(search)
}