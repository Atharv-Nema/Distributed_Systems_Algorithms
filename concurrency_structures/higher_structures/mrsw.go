package Structures

import "sync"

type MRSW[T any] struct {
	readlock  sync.Mutex
	writelock sync.Mutex
	turn      sync.Mutex
	nr        int
	data      T
}

func (m *MRSW[T]) ReadData() T {
	/* Readlock is only present so that m.nr can be accessed and modified in
	   a thread safe manner.
	   Note the general pattern observed.
	   To implement a Semaphore, we needed to use a Spinlock for the internal
	   variables
	   To implement an MRSW, we needed a single reader/single writer for nr.
	   The idea is that this is way less computationally expensive that the
	   read of the actual data, hence this is a reasonable thing to do
	   If we now remove all the m.turn lines, we still have a working solution.
	   But, notice how the writelock is only released after all readers are done.
	   If there are many readers, the writer can starve.
	   To solve this, we use another lock called turn. For any entity(reader or writer)
	   to start its task, it must get hold of this lock. Every entity(reader or writer)
	   will eventually get hold of this lock(because of the queue datastructure). When a
	   writer gets this lock, it releases it only after it has acquired the write lock(so it
	   will eventually get the write lock). When a reader acquires the turn lock, this means
	   that it was infront of any writer in the queue(it has fairly reached there). As a reader
	   may not get hold of the writelock(another reader may have acquired it), you may think that
	   it must release the turn after it gets the readlock. That is a valid thing to do. However,
	   note that the only real lock is the writelock(the readlock is just needed to make accesses
	   to nr thread safe). So, any fairness in the readlock is not really needed(i.e. it is fine
	   if a reader behind 'this' reader in the turn queue acquires the readlock before 'this'.
	   Hence the strange empty critical section)
	*/

	m.turn.Lock()
	m.turn.Unlock()
	m.readlock.Lock()
	m.nr++
	if m.nr == 1 {
		// First reader
		m.writelock.Lock()
	}
	m.readlock.Unlock()
	readData := m.data
	m.readlock.Lock()
	m.nr--
	if m.nr == 0 {
		// Last reader
		m.writelock.Unlock()
	}
	m.readlock.Unlock()
	return readData
}

func (m *MRSW[T]) WriteData(newData T) {
	m.turn.Lock()
	m.writelock.Lock()
	m.turn.Unlock()
	m.data = newData
	m.writelock.Unlock()
}
