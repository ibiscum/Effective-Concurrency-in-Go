package main

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}

func main() {
	for _, s := range []string{"a", "b", "c", "d", "e"} {
		go f(s)
	}
	time.Sleep(2 * time.Second)
}
