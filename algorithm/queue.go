package algorithm

// NewQueue create new queue
func NewQueue(length int) *Queue {
	return &Queue{Length: length}
}

func (q *Queue) IsFull() bool {
	if q.Rear == q.Length {
		// fmt.Println("Queue is full")
		return true
	}
	return false
}

func (q *Queue) IsEmpty() bool {
	if q.Front == q.Rear {
		// fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (q *Queue) enQueue(n *BTree) {
	if q.IsFull() {
		return
	}
	// fmt.Println("in queue ", n.Data)
	q.Elements = append(q.Elements[:q.Rear], n)
	q.Rear++
}

func (q *Queue) deQueue() *BTree {
	if q.IsEmpty() {
		return nil
	}
	// fmt.Println("out queue ", q.Elements[q.Front].Data)
	out := q.Elements[q.Front]
	q.Elements = q.Elements[q.Front+1 : q.Rear]
	q.Rear--
	return out
}
