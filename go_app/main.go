// Created by: Jaden Plugowsky
// Created on: April 2023
//
// This program converts Fahrenheit temperature into Celsius

package main

import "fmt"

func main() {
	// This function converts Fahrenheit temperature into Celsius
	var fahrenheit float64
	var celsius float64

	// Input
	fmt.Println("\nThis program converts an inputted Fahrenheit temperature value into Celsius.")
	fmt.Println()
	fmt.Print("Enter the temperature in Fahrenheit (°F): ")
	fmt.Scanln(&fahrenheit)

	// Process
	celsius = (fahrenheit - 32) * 5 / 9

	// Output
	fmt.Printf("\nThe temperature in Celsius is: %.2f °C.", celsius)

	fmt.Println("\n\nDone.")
}
