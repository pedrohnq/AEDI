package lists

import (
	"errors"
)

type ArrayList struct {
	values []int
	size   int
}

func (arraylist *ArrayList) Init(size int) error {
	if size > 0 {
		arraylist.values = make([]int, size)
		return nil
	} else {
		return errors.New("size must be greater than 0")
	}
}

func (arraylist *ArrayList) double() {
	newSize := arraylist.size * 2
	newValues := make([]int, newSize)
	for i := 0; i < arraylist.size; i++ {
		newValues[i] = arraylist.values[i]
	}
	arraylist.values = newValues
}

func (arraylist *ArrayList) Add(value int) error {
	if arraylist.size == len(arraylist.values) {
		arraylist.double()
	}
	arraylist.values[arraylist.size] = value
	arraylist.size++
	return nil
}

func (arraylist *ArrayList) AddOnIndex(value, index int) error {
	if index >= 0 && index <= arraylist.size {
		if arraylist.size == len(arraylist.values) {
			arraylist.double()
		}
		for i := arraylist.size; i > index; i-- {
			arraylist.values[i] = arraylist.values[i-1] // shift right
		}
		arraylist.values[index] = value
		arraylist.size++
		return nil
	} else {
		return errors.New("index out of bounds")
	}
}

func (arraylist *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < arraylist.size {
		for i := index; i < arraylist.size-1; i++ {
			arraylist.values[i] = arraylist.values[i+1] // shift left
		}
		arraylist.values[arraylist.size-1] = 0
		arraylist.size--
		return nil
	} else {
		return errors.New("index out of bounds")
	}
}

func (arraylist *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < arraylist.size {
		return arraylist.values[index], nil
	} else {
		return 0, errors.New("index out of bounds")
	}
}

func (arraylist *ArrayList) Set(value, index int) error {
	if index >= 0 && index < arraylist.size {
		arraylist.values[index] = value
		return nil
	} else {
		return errors.New("index out of bounds")
	}
}

func (arraylist *ArrayList) Size() int {
	return arraylist.size
}

func (arraylist *ArrayList) Invert() {
	if arraylist.size > 0 {
		new_list := make([]int, len(arraylist.values))
		idx := 0
		for i := arraylist.size - 1; i >= 0; i-- {
			new_list[idx] = arraylist.values[i]
			idx++
		}
		arraylist.values = new_list
	}
}
