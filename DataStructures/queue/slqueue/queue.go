package slqueue

type DataType interface{}

type Queue []DataType

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Enqueue(d DataType) {
	*q = append((*q), d)
}

func (q *Queue) Peek() DataType {
	return (*q)[0]
}

func (q *Queue) Dequeue() DataType {
	if q.IsEmpty() {
		return nil
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}
