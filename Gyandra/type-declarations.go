package main 

import "fmt"

func main () {
	type NoKTP string
	type Married bool

	var NoKtpGya NoKTP = "1341561715"
	var marriedStatus Married = true
	fmt.Println(NoKtpGya)
	fmt.Println(marriedStatus)

}
