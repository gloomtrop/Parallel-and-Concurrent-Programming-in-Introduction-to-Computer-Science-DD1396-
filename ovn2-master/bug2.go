package main

import (
    "fmt"
    "sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
// The bug is that the main is closed before the goroutine is finnished
//Therefore a WaitGroup is added to wait for all the goroutines to execute
func main() {
    ch := make(chan int)
    wait := new(sync.WaitGroup)
    wait.Add(1)
    go Print(ch, wait)
    for i := 1; i <= 11; i++ {
        ch <- i
    }
    wait.Wait()
    close(ch)
}
// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup){
    for n := range ch { // reads from channel until it's closed
        fmt.Println(n)
    }
    wg.Done()
}
