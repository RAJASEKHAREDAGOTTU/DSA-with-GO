package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")

		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}

	}
	fmt.Println()
}
func main() {

	FizzBuzz(20)
}