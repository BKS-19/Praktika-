package main

import "fmt"

func main() {
	var a [5]int

	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}

	sum := 0
	for i := 0; i < 5; i++ {
		sum += a[i]
	}
	fmt.Println(sum)

	max := a[0]
	maxIndex := 0
	for i := 1; i < 5; i++ {
		if a[i] > max {
			max = a[i]
			maxIndex = i
		}
	}
	fmt.Println(max, maxIndex)

	positive := make([]int, 0)
	for i := 0; i < 5; i++ {
		if a[i] > 0 {
			positive = append(positive, a[i])
		}
	}

	if len(positive) > 0 {
		for i, val := range positive {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	} else {
		fmt.Println("")
	}
}
