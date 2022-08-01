package main

import (
	"fmt"
)

func main() {
	/*
		Print numbers between 1 to 20
		- If the number is divisible by 3 print "fizz"
		- If the number is divisible by 5 print "buzz"
		- If the number is divisible by 3 and 5  print "fizz buzz"
	*/

	for i := 1; i <= 20; i++ {
		value := i
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(value)
		}
	}
}
