package main

import "fmt"

func factoriaLLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factoriaLRecursive(value int) int {
	if value == 1 {
		return 1
	}else {
		return value * factoriaLRecursive(value-1)
	}
}

func main() {
	loop := factoriaLLoop(5)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factoriaLRecursive(5)
	fmt.Println(recursive)
}

