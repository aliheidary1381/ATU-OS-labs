package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id               int
	leftChopstick    *Chopstick
	rightChopstick   *Chopstick
	arbitrator       *sync.Mutex
	eatingCounter    int
	maxEatingCounter int
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * time.Duration(500))
}

func (p *Philosopher) eat() {
	p.arbitrator.Lock()
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()

	fmt.Printf("Philosopher %d is eating\n", p.id)
	p.eatingCounter++
	time.Sleep(time.Millisecond * time.Duration(500))

	p.rightChopstick.Unlock()
	p.leftChopstick.Unlock()
	p.arbitrator.Unlock()
}

func (p *Philosopher) dine() {
	for p.eatingCounter < p.maxEatingCounter {
		p.think()
		p.eat()
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{}
	}

	arbitrator := &sync.Mutex{}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:               i + 1,
			leftChopstick:    chopsticks[i],
			rightChopstick:   chopsticks[(i+1)%numPhilosophers],
			arbitrator:       arbitrator,
			maxEatingCounter: 3, 
		}
	}

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(p *Philosopher) {
			defer wg.Done()
			p.dine()
		}(philosopher)
	}

	wg.Wait()
	fmt.Println("Dinner is over!")
}
