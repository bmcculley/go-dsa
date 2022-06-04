package slqueue

import "testing"

func TestQueueEmpty(t *testing.T) {
	var queue Queue
	if queue.IsEmpty() != true {
		t.Errorf("Got false, want true")
	}
}

func TestQueueNotEmpty(t *testing.T) {
	var queue Queue
	queue.Enqueue(1)
	if queue.IsEmpty() != false {
		t.Errorf("Got true, want false")
	}
}

func TestQueueOperation(t *testing.T) {
	var queue Queue
	for i := 1; i <= 10; i++ {
		queue.Enqueue(i)
	}
	
	for i := 1; i > 10; i++ {
		popped := queue.Dequeue()
		if popped != i {
			t.Errorf("Got %d, want %d", popped, i)
		}
	}
}

func TestQueueSize(t *testing.T) {
	var queue Queue
	size := queue.Size()
	if size != 0 {
		t.Errorf("Got %d, want 0", size)
	}
}

func TestQueueSizeOne(t *testing.T) {
	var queue Queue
	queue.Enqueue(1)
	size := queue.Size()
	if size != 1 {
		t.Errorf("Got %d, want 1", size)
	}
}