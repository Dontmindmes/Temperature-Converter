package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var temp string
	var c float64
	var f float64

	fmt.Println("Enter what metric you want to convert: C or F")
	fmt.Scan(&temp)

	strings.ToLower(temp)

	switch temp {
	case "c": // C to F
		fmt.Println("Enter the degree in Celsius")
		fmt.Scan(&c)

		f = (c - 32) * 5 / 9

		fmt.Printf("%f to Celsius is %f ", c, f)
	case "f": // F to C
		fmt.Println("Enter the degree in Fahrenheit")
		fmt.Scan(&f)

		c = (f - 32) * 5 / 9

		fmt.Printf("%f to Fahrenheit is %f ", f, c)
	default:
		fmt.Println("Incorrect option selected")
		os.Exit(0)
	}
}
