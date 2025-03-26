package Structures

// This is kind of how go channels work. The only difference is
// that we cannot have zero buffering

import (
	DS "concurrency_structures/helper_datastructures"
	Lock "concurrency_structures/locks"
	"sync"
)

type Channel[T any] struct {
	spaces *Lock.Semaphore
	items  *Lock.Semaphore
	buffer *DS.CircularBuffer[T]
	guard  sync.Mutex
}

func (c *Channel[T]) BufferedItemCount() uint32 {
	return c.buffer.NumItems()
}

func NewChannel[T any](size uint32) *Channel[T] {
	channel := new(Channel[T])
	channel.items = Lock.NewSemaphore(0)
	channel.spaces = Lock.NewSemaphore(size)
	channel.buffer = DS.NewCircularBuffer[T](size)
	return channel
}

func (c *Channel[T]) Produce(pid int, product T) {
	c.spaces.Wait(pid) // Waiting until spaces are available
	c.guard.Lock()
	c.buffer.AddItem(product) // Not thread safe
	c.guard.Unlock()
	c.items.Signal(pid)
}

func (c *Channel[T]) Consume(pid int) T {
	c.items.Wait(pid)
	c.guard.Lock()
	product := c.buffer.PopItem()
	c.guard.Unlock()
	c.spaces.Signal(pid)
	return product
}
