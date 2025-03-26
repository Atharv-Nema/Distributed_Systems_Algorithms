package main

import (
	// Lock "concurrency_structures/locks"
	Simul "concurrency_structures/simulations"
)

func main() {
	// fridgeLock := Lock.NewMutex()
	// Simul.RunDrunkardSimulation(fridgeLock)
	Simul.RunProducerConsumerSimulation(1000, 10, 20)
}
