package Lock

// Lock interface that both SpinLock and LamportLock implement
type Lock interface {
	Lock(pid int)
	Unlock(pid int)
}
