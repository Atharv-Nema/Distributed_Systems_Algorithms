package main

import (
	Lock "concurrency_structures/locks"
)

func main() {
	fridgeLock := Lock.NewMutex()
	Lock.RunSimulation(fridgeLock)
}
