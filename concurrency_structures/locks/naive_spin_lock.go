package Lock

import (
	"sync/atomic"
)

type SpinLock uint32 // This is uint32 not a simple bool because the CAS function expects a uint32

func NewSpinLock() *SpinLock {
	lock := SpinLock(0)
	return &lock
}

// Lock function
func (sl *SpinLock) Lock(_ int) {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		continue
	}
}

// Unlock function
func (sl *SpinLock) Unlock(_ int) {
	*sl = 0
}
