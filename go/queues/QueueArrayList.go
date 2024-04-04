package queues

import (
	"aedi/go/lists"
	"errors"
)

type QueueArrayList struct {
	data lists.ArrayList
	size int
}

func (queue *QueueArrayList) Enqueue(value int) {
	if queue.size == 0 {
		queue.data.Init(1)
	}
	queue.data.Add(value)
	queue.size++
}

func (queue *QueueArrayList) Dequeue() (int, error) {
	if queue.size == 0 {
		return 0, errors.New("queue is empty")
	}
	value, _ := queue.Peek()
	queue.data.RemoveOnIndex(0)
	queue.size--
	return value, nil
}

func (queue *QueueArrayList) Peek() (int, error) {
	if queue.size == 0 {
		return 0, errors.New("queue is empty")
	}
	value, _ := queue.data.Get(0)
	return value, nil
}

func (queue *QueueArrayList) Size() int {
	return queue.size
}
