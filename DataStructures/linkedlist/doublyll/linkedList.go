package doublyll

import "fmt"

type DataType interface{}

type node struct {
	data DataType
	next *node
	prev *node
}

type LinkedList struct {
	head *node
	tail *node
	length int
}

func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList) Size() int {
	return l.length
}

func (l *LinkedList) Add(d DataType) {
	l.length++
	if l.head == nil {
		l.head = &node{d, nil, nil}
		l.tail = l.head
	} else {
		temp := l.head
		prev := temp
		for temp.next != nil {
			prev = temp
			temp = temp.next
		}
		temp.next = &node{d, nil, nil}
		temp.next.prev = prev
		l.tail = temp
	}
}

func (l *LinkedList) Push(d DataType) {
	l.length++
	if l.head == nil {
		l.head = &node{d, nil, nil}
		l.tail = l.head
	} else {
		temp := &node{d, nil, nil}
		temp.next = l.head
		l.head.prev = temp
		l.head = temp
	}
}

func (l *LinkedList) Pop() DataType {
	if l.IsEmpty() {
		return nil
	} else {
		l.length--
		temp := l.head.data
		if l.head.next != nil {
			l.head = l.head.next
		}
		l.head.prev = nil
		return temp
	}
}

func (l *LinkedList) Reverse() {
	head := l.head
	var prev *node
	var curr *node
	
	keepGoing := true

	for keepGoing {
		if head.next == nil {
			keepGoing = false
		}
		curr = head
		head = curr.next
		curr.next = prev
		curr.prev = head
		prev = curr
	}
	l.head = prev
}

func (l *LinkedList) PrintList() {
	temp := l.head
	keepGoing := true

	for keepGoing {
		fmt.Println(temp.data)
		if temp.next != nil {
			temp = temp.next
		} else {
			keepGoing = false
		}
	}
}
