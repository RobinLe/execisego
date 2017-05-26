package algorithm

import "fmt"

// Queue first in first out FIFO
type Queue struct {
	Elements []*Element
	Rear     int
	Front    int
	Length   int
}

// NewQueue create new queue
func NewQueue(length int) *Queue {
	return &Queue{Length: length}
}

func (q *Queue) isFull() bool {
	if q.Rear == q.Length {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (q *Queue) isEmpty() bool {
	if q.Front == q.Rear {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (q *Queue) enQueue(n *Element) {
	if q.isFull() {
		return
	}
	fmt.Println("in queue ", n.Value)
	q.Elements = append(q.Elements[:q.Rear], n)
	q.Rear++
}

func (q *Queue) deQueue() {
	if q.isEmpty() {
		return
	}
	fmt.Println("out queue ", q.Elements[q.Front].Value)
	q.Elements = q.Elements[q.Front+1 : q.Rear]
	q.Rear--
}

func QueueTest() {
	q := NewQueue(3)
	q.deQueue()
	q.enQueue(&Element{1})
	q.enQueue(&Element{2})
	q.enQueue(&Element{3})
	q.enQueue(&Element{4})
	q.deQueue()
	q.deQueue()
	q.deQueue()
	q.deQueue()
	q.enQueue(&Element{4})
	q.enQueue(&Element{5})
	q.deQueue()
	q.deQueue()
}
