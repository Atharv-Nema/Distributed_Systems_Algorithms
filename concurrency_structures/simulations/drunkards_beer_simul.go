package Simulations

import (
	Lock "concurrency_structures/locks"
	"fmt"
	"math/rand"
	"time"
)

var (
	numBeers int
)

// Simulating the beer buying problem
func consumeBeer(id int, fridgeLock Lock.Lock) {
	fridgeLock.Lock(id)
	if numBeers > 0 {
		numBeers--
		fmt.Printf("Process id %d consumed beer\n", id)
		fmt.Printf("Remaining beers are %d\n", numBeers)
	}
	fridgeLock.Unlock(id)
}

func buyBeer(id int, fridgeLock Lock.Lock) {
	fridgeLock.Lock(id)
	if numBeers == 0 {
		numBeers += 5
		fmt.Printf("Process id %d bought beer\n", id)
		fmt.Printf("Remaining beers are %d\n", numBeers)
	}
	fridgeLock.Unlock(id)
}

func alcoholic(id int, fridgeLock Lock.Lock) {
	for {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		drinking_choice := rand.Intn(10)
		if drinking_choice < 2 {
			buyBeer(id, fridgeLock)
		} else {
			consumeBeer(id, fridgeLock)
		}
	}
}

func RunDrunkardSimulation(fridgeLock Lock.Lock) {
	numBeers = 0
	for i := range 10 {
		go alcoholic(i, fridgeLock)
	}
	var input string
	fmt.Scanln(&input)
}
