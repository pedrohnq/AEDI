package stacks

type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	Size() int
}
