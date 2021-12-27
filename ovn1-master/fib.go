package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	num0 := 0
	num1 := 1
	return func(x int) int {
		switch x {
		case 0:
			return num0
		case 1:
			return num1
		}
		num2 := num0 + num1
		num0 = num1
		num1 = num2
		return num2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
