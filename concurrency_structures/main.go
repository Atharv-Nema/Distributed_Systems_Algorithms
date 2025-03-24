package main

import (
	Lock "concurrency_structures/locks"
)

func main() {
	fridgeLock := new(Lock.LamportLock)
	Lock.RunSimulation(fridgeLock)
}
