package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	s = "a"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	s = "b"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	s = "c"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	time.Sleep(1 * time.Second)
}
