package main

import "fmt"

func main() {
	ch := make(chan bool)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(i)
		}
	}()

	go func() {
		fmt.Println("hello")
	}()

	<-ch
}
