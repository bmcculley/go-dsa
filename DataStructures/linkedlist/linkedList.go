package linkedlist

import "fmt"

type DataType interface{}

type node struct {
	data DataType
	next *node
}

type LinkedList struct {
	head *node
	length int
}

func (l *LinkedList) Add(d DataType) {
	l.length++
	if l.head == nil {
		l.head = &node{d, nil}
	} else {
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &node{d, nil}
	}
}

func (l *LinkedList) Push(d DataType) {
	l.length++
	if l.head == nil {
		l.head = &node{d, nil}
	} else {
		temp := &node{d, nil}
		temp.next = l.head
		l.head = temp
	}
}

func (l *LinkedList) Pop() DataType {
	if l.IsEmpty() {
		return ""
	} else {
		l.length--
		temp := l.head.data
		l.head = l.head.next
		return temp
	}
}

func (l *LinkedList) Reverse() {
	head := l.head.next
	temp := l.head.next
	prev := l.head
	prev.next = nil
	keepGoing := true

	for keepGoing {
		temp = head
		head = head.next
		temp.next = prev
		prev = temp
		if head.next == nil {
			keepGoing = false
		}
	}
	head.next = temp
	l.head = head
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

func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l *LinkedList) Size() int {
	return l.length
}
