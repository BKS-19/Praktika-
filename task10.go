package main

import "fmt"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println("division by zero")
		return
	}

	quotient, remainder := divmod(a, b)

	fmt.Println(quotient, remainder)
}
