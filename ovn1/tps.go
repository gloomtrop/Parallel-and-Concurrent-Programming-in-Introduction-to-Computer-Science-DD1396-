package main

import "fmt"

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
	// TODO
	sum := 0
	for _,v := range a {
		sum += v
	}
	res <- sum
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    go Add(a[:n/2], ch)
    go Add(a[n/2:], ch)

	// TODO: Get the subtotals from the channel and print their sum.
	part1, part2 := <-ch, <-ch
	sum := part1 + part2
	fmt.Println(sum)

	// Test 
	sum1 := 0
	for _,v := range a {
		sum1 += v
	}
	fmt.Println(sum1)
}