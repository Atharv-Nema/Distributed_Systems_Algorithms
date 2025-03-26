package DS

type CircularBuffer[T any] struct {
	head     uint32
	tail     uint32
	size     uint32
	capacity uint32
	buffer   []T
}

func (c *CircularBuffer[T]) NumItems() uint32 {
	return c.size
}

func NewCircularBuffer[T any](N uint32) *CircularBuffer[T] {
	if N < 1 {
		panic("Buffer size must be >= 0")
	}
	circularBuffer := new(CircularBuffer[T])
	circularBuffer.buffer = make([]T, N)
	circularBuffer.head = 0
	circularBuffer.tail = 0
	circularBuffer.capacity = N
	return circularBuffer
}

func (c *CircularBuffer[T]) AddItem(item T) {
	if c.size == c.capacity {
		panic("Buffer is full")
	}
	c.buffer[c.head] = item
	c.head++
	c.head %= c.capacity
	c.size++
}

func (c *CircularBuffer[T]) PopItem() T {
	if c.size == 0 {
		panic("Buffer is empty")
	}
	item := c.buffer[c.tail]
	c.tail++
	c.tail %= c.capacity
	c.size--
	return item
}
