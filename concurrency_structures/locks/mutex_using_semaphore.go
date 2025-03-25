package Lock

type Mutex struct {
	semaphore  *Semaphore
	simpleLock *SpinLock
}

func NewMutex() *Mutex {
	mutex := new(Mutex)
	mutex.simpleLock = NewSpinLock()
	mutex.semaphore = NewSemaphore(1)
	return mutex
}

func (m *Mutex) Lock(pid int) {
	// Precondition: self.semaphore.capacity <= 1
	m.semaphore.Wait(pid)
}

func (m *Mutex) Unlock(pid int) {
	m.simpleLock.Lock(pid)
	if m.semaphore.current_capacity == 0 {
		m.semaphore.Signal(pid)
	}
	m.simpleLock.Unlock(pid)
}
