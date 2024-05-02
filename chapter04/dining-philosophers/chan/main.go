package main

import (
	"fmt"
	"math/rand"
	"time"
)

func philosopher(index int, leftFork, rightFork chan bool) {
	for {
		// Think for some time
		fmt.Printf("Philosopher %d is thinking\n", index)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		select {
		case <-leftFork:
			leftFork <- true
			select {
			case <-rightFork:
				rightFork <- true
				fmt.Printf("Philosopher %d is eating\n", index)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			default:
			}
		default:
		}
	}
}

func main() {
	var forks [5]chan bool
	for i := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	go philosopher(0, forks[4], forks[0])
	go philosopher(1, forks[0], forks[1])
	go philosopher(2, forks[1], forks[2])
	go philosopher(3, forks[2], forks[3])
	go philosopher(4, forks[3], forks[4])
	select {}
}
