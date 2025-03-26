package DS

type dequeNode[T any] struct {
	value T
	next  *dequeNode[T]
	prev  *dequeNode[T]
}

type Deque[T any] struct {
	head *dequeNode[T]
	tail *dequeNode[T]
}

func NewDeque[T any]() *Deque[T] {
	deque := new(Deque[T])
	deque.head = nil
	deque.tail = nil
	return deque
}

func (self *Deque[T]) AddHead(item T) {
	newNode := new(dequeNode[T])
	newNode.value = item
	newNode.next = nil
	newNode.prev = nil
	if self.head == nil && self.tail == nil {
		self.head = newNode
		self.tail = newNode
	} else {
		newNode.next = self.head
		self.head.prev = newNode
		self.head = newNode
	}
}

func (self *Deque[T]) PopTail() T {
	if self.tail == nil {
		panic("List is empty")
	}
	toReturn := self.tail.value
	if self.tail.prev == nil {
		self.head = nil
		self.tail = nil
	} else {
		self.tail = self.tail.prev
		self.tail.next = nil
	}
	return toReturn
}
