package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z, zold float64 = 1.0, 0.0
	for i := 0; i<10; i++{
		z -= (z*z - x) / (2*z)
		if math.Abs(z-zold) <math.Pow(10,-10) {
			i = 10
		}
		zold = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
