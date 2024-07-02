package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	go func() {
		for {
			mutex.Lock()
			// Work for a long time
			fmt.Println("go 1")
			mutex.Unlock()
			// Do something quick
			os.Exit(1)
		}
	}()

	go func() {
		for {
			mutex.Lock()
			// Do something quick
			fmt.Println("go 2")
			mutex.Unlock()
			// Work for a long time
			os.Exit(2)
		}
	}()

	time.Sleep(2 * time.Second)

}
