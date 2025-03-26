package Simulations

import (
	ST "concurrency_structures/higher_structures"
	"fmt"
	"math/rand"
	"time"
)

func Producer(pid int, c *ST.Channel[int]) {
	for {
		item := rand.Intn(1000) + 1000
		productionTime := time.Millisecond * time.Duration(rand.Intn(250))
		time.Sleep(productionTime)
		fmt.Printf("Producer with pid %d produced %d, items in buffer: %d\n", pid, item, c.BufferedItemCount())
		c.Produce(pid, item)
	}
}

func Consumer(pid int, c *ST.Channel[int]) {
	for {
		item := c.Consume(pid)
		fmt.Printf("Consumer with pid %d consumed %d, items in buffer: %d\n", pid, item, c.BufferedItemCount())
	}
}

func RunProducerConsumerSimulation(numProducer int, numConsumer int, chanSize uint32) {
	c := ST.NewChannel[int](uint32(chanSize))
	for i := range numProducer {
		go Producer(i, c)
	}
	for i := range numConsumer {
		go Consumer(i+numProducer, c)
	}
	var input string
	fmt.Scanln(&input)
}
