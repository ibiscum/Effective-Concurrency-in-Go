package main

import (
	"fmt"
	"time"
)

func main() {
	for _, s := range []string{"a", "b", "c"} {
		s := s

		go func() {
			fmt.Printf("Goroutine %s\n", s)
		}()
	}
	time.Sleep(1 * time.Second)
}
