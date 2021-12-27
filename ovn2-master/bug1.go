package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
    ch := make(chan string, 1) //Noone is recieving which means it is blocked --> make buffer
    ch <- "Hello world!"
    fmt.Println(<-ch)
}