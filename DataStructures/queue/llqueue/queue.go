package slqueue

type DataType interface{}

type node struct {
	data DataType
	next *node
}

type Queue struct {
	head *node
	tail *node
	length int
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) Enqueue(d DataType) {
	q.length++
	if q.head == nil {
		q.head = &node{d, nil}
		q.tail = q.head
	} else {
		temp := &node{d, nil}
		q.tail.next = temp
		q.tail = temp
	}
}

func (q *Queue) Peek() DataType {
	return q.head.data
}

func (q *Queue) Dequeue() DataType {
	if q.IsEmpty() {
		return nil
	} else {
		q.length--
		temp := q.head.data
		q.head = q.head.next
		return temp
	}
}
