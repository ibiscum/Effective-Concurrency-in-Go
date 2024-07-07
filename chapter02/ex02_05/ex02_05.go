package main

import (
	"fmt"
	"time"
)

func main() {
	chn := make(chan bool) // Create an unbuffered channel
	go func() {
		chn <- true // Send to channel
	}()

	go func() {
		y := <-chn // Receive from channel
		fmt.Println(y)
	}()

	time.Sleep(1 * time.Second)
}
