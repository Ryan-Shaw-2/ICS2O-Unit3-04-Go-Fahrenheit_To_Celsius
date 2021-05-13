// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program coverts fahrenheit to celsius

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function converts fahrenheit to celsius
	var fahrenheit float64

	// input
	fmt.Println("This program converts fahrenheit to celsius")
	fmt.Println()
	fmt.Print("Enter the temperature in fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	fmt.Println()

	// process
	var celcius = (fahrenheit - 32) * (5.0/9.0)

	// output
	fmt.Println(fahrenheit, "°F is", math.Round(celcius*1000)/1000, "°C")
}
