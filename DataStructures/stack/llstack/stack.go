package llstack

type DataType interface{}

type node struct {
	data DataType
	next *node
}

type Stack struct {
	head *node
	length int
}

func (l *Stack) Size() int {
	return l.length
}

func (l *Stack) IsEmpty() bool {
	return l.length == 0
}

func (l *Stack) Push(d DataType) {
	l.length++
	if l.head == nil {
		l.head = &node{d, nil}
	} else {
		temp := &node{d, nil}
		temp.next = l.head
		l.head = temp
	}
}

func (l *Stack) Pop() DataType {
	if l.IsEmpty() {
		return ""
	} else {
		l.length--
		temp := l.head.data
		l.head = l.head.next
		return temp
	}
}
