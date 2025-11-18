package main

import "fmt"

func bubbleSort(a *[5]int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {

				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	var a [5]int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}

	bubbleSort(&a)

	for i := 0; i < 5; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}
