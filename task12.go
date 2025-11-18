package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	var temp float64

	fmt.Scan(&temp)

	celsiusTemp := Celsius(temp)

	fahrenheit := celsiusTemp.ToFahrenheit()
	fmt.Printf("%.2f\n", fahrenheit)
}
