package queue

type QueueImp[Type any] interface {
	Enqueue(value Type)
	Dequeue() Type
	IsEmpty() bool
}
