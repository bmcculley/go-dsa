package llstack

import "testing"

func TestStackEmpty(t *testing.T) {
	var stack Stack
	if stack.IsEmpty() != true {
		t.Errorf("Got false, want true")
	}
}

func TestStackNotEmpty(t *testing.T) {
	var stack Stack
	stack.Push(1)
	if stack.IsEmpty() != false {
		t.Errorf("Got true, want false")
	}
}

func TestStackOperation(t *testing.T) {
	var stack Stack
	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}
	
	for i := 10; i > 0; i-- {
		popped := stack.Pop()
		if popped != i {
			t.Errorf("Got %d, want %d", popped, i)
		}
	}
}

func TestStackSize(t *testing.T) {
	var stack Stack
	size := stack.Size()
	if size != 0 {
		t.Errorf("Got %d, want 0", size)
	}
}

func TestStackSizeOne(t *testing.T) {
	var stack Stack
	stack.Push(1)
	size := stack.Size()
	if size != 1 {
		t.Errorf("Got %d, want 1", size)
	}
}