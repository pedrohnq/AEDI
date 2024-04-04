package queues

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Peek() (int, error)
	Size() int
}
