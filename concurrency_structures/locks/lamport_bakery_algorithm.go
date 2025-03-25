package Lock

type LamportLock struct {
	Enter  [10]bool
	Number [10]int
}

// The basic idea of the algorithm is to first assign ticket numbers, and then call on those numbers.
// The waiting queues just choose a number larger than all the ones present. The point of Enter is so that
// if there is a process making no requests ever, we do not wait unnecessarily for it. We need to wait for
// a process to be assigned a ticket before performing the order check because it may be possible for the process
// to have calculated the max_value, but be stuck just before assigning it. Enter[pid] is used to prevent these cases,
// as we wait until all things have been assigned.

// Question: How are array reads safe here?

func (ll *LamportLock) Lock(pid int) {
	ll.Enter[pid] = true
	max_val := ll.Number[0]
	for _, value := range ll.Number {
		if value > max_val {
			max_val = value
		}
	}
	ll.Number[pid] = max_val + 1
	ll.Enter[pid] = false

	for j := 0; j < 10; j++ {
		for ll.Enter[j] {
			continue
			// If a process has already reached the Enter[pid] stage, we wait until it has been assigned a ticket id
		}
		for (ll.Number[j] > 0) && (ll.Number[j] < ll.Number[pid] || (ll.Number[j] < ll.Number[pid] && j < pid)) {
			// Number[j] != 0 means that it is waiting on the lock
			// It has a higher priority that this process if the Number[j] ... condition holds
			// If so, we wait until it completes(note: it may not have the lock yet, it may just be waiting for it)
			continue
		}
	}
}

func (ll *LamportLock) Unlock(pid int) {
	ll.Number[pid] = 0
}
