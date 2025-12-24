# Implementation of the queue data structure

To create a queue, use the command:

```golang

Queue := NewQueue[Type]() // [Type] -> Defines what type of data the queue will store.

```

If you want a queue of type string:

```golang
Queue := NewQueue[string]()
```

To enqueue, use the command:

```golang
Queue.Enqueue("<The type that corresponds to what you chose for the queue.>")
// "The type..."
// first in

Queue.Enqueue("Other value")
// "The type...", ["Other value"]
// first in,        Two in
```

To dequeue, use:

```golang

value := Queue.Dequeue()
// x, "Other value"
// first out, now the first element in the queue is the second    [Two in]

```
