package Lock

import (
	DS "concurrency_structures/helper_datastructures"
)

// Unfortunately, go does not allow us to halt and resume individual goroutines directly.
// Hence, I am implementing semaphores using a queue of channels, which honestly is pretty stupid
// (we should be implementing channels with mutex's)

type Semaphore struct {
	outerlock        *SpinLock
	current_capacity uint32
	queue            *DS.Deque[chan int]
	waiting_ct       int
}

func NewSemaphore(initial_capacity uint32) *Semaphore {
	s := new(Semaphore)
	s.current_capacity = initial_capacity
	s.queue = DS.NewDeque[chan int]()
	s.waiting_ct = 0
	s.outerlock = new(SpinLock)
	return s
}

func (s *Semaphore) Wait(pid int) {
	s.outerlock.Lock(pid)
	if s.current_capacity > 0 {
		s.current_capacity--
		s.outerlock.Unlock(pid)
	} else {
		s.waiting_ct++
		newChannel := make(chan int)
		s.queue.AddHead(newChannel)
		s.outerlock.Unlock(pid)
		newChannel <- 0 // This adds the process to the queue
	}
}

func (s *Semaphore) Signal(pid int) {
	s.outerlock.Lock(pid)
	if s.waiting_ct == 0 {
		s.current_capacity++
		s.outerlock.Unlock(pid)
	} else {
		s.waiting_ct--
		s.outerlock.Unlock(pid)
		<-s.queue.PopTail()
	}
}
