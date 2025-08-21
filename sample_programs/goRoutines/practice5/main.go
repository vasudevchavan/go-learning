package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numPhilosophers = 5

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

func (p Philosopher) dine(wg *sync.WaitGroup, host chan struct{}) {
	defer wg.Done()
	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		// Think
		fmt.Printf("Philosopher %d is thinking\n", p.id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

		// Request to eat (hosted)
		host <- struct{}{} // Request permission from host

		// Pick up forks (to prevent deadlock, pick in order)
		if p.id%2 == 0 {
			p.leftFork.Lock()
			p.rightFork.Lock()
		} else {
			p.rightFork.Lock()
			p.leftFork.Lock()
		}

		// Eat
		fmt.Printf("Philosopher %d is eating\n", p.id)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

		// Put down forks
		p.leftFork.Unlock()
		p.rightFork.Unlock()

		<-host // Done eating, release slot from host
	}
	fmt.Printf("Philosopher %d is done eating\n", p.id)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	forks := make([]*Fork, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &Fork{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	var wg sync.WaitGroup
	host := make(chan struct{}, 2) // Only 2 philosophers can eat at a time

	for _, p := range philosophers {
		wg.Add(1)
		go p.dine(&wg, host)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating.")
}
