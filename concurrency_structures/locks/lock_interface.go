package Lock

// Lock interface that both SpinLock and LamportLock implement
type Lock interface {
	lock(pid int)
	unlock(pid int)
}
