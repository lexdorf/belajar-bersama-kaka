package main

import(
	"errors"
	"fmt"
)

func Pembagi(nilai int, Pembagi int) (int, error) {
	if Pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}else {
		result := nilai / Pembagi
		return result, nil
	}
}

func main() {
	hasil, err := Pembagi(100,0)
	if err == nil{
		fmt.Println("Hasil", hasil)
	}else {
		fmt.Println("Error", err.Error())
	}

}