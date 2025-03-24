package Lock

import (
	"sync/atomic"
)

type SpinLock uint32 // This is uint32 not a simple bool because the CAS function expects a uint32

// Lock function
func (sl *SpinLock) lock(_ int) {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		continue
	}
}

// Unlock function
func (sl *SpinLock) unlock(_ int) {
	*sl = 0
}
