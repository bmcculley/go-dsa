package singlyll

import "testing"

func TestLinkedListEmpty(t *testing.T) {
	var linkedlist LinkedList
	if linkedlist.IsEmpty() != true {
		t.Errorf("Got false, want true")
	}
}

func TestLinkedListNotEmpty(t *testing.T) {
	var linkedlist LinkedList
	linkedlist.Push(1)
	if linkedlist.IsEmpty() != false {
		t.Errorf("Got true, want false")
	}
}

func TestLinkedListOperation(t *testing.T) {
	var linkedlist LinkedList
	for i := 1; i <= 10; i++ {
		linkedlist.Push(i)
	}
	
	for i := 10; i > 0; i-- {
		popped := linkedlist.Pop()
		if popped != i {
			t.Errorf("Got %d, want %d", popped, i)
		}
	}
}

func TestLinkedListSize(t *testing.T) {
	var linkedlist LinkedList
	size := linkedlist.Size()
	if size != 0 {
		t.Errorf("Got %d, want 0", size)
	}
}

func TestLinkedListSizeOne(t *testing.T) {
	var linkedlist LinkedList
	linkedlist.Push(1)
	size := linkedlist.Size()
	if size != 1 {
		t.Errorf("Got %d, want 1", size)
	}
}