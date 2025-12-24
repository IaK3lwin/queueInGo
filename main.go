package queue

import (
	"container/list"
)


type Queue[Type any] struct {
	data *list.List
}

func NewQueue[Type any]() *Queue[Type] {
	return &Queue[Type]{
		data: list.New(),
	}
}

//Enqueue
func (q *Queue[Type]) Enqueue(value Type) {
	q.data.PushBack(value)
}

func (q *Queue[Type]) Dequeue() Type {

	value := q.data.Front()

	q.data.Remove(value)

	return value.Value.(Type)
}

func (q *Queue[Type]) IsEmpty() bool {
	if q.data.Len() == 0 {
		return true
	}

	return  false
}