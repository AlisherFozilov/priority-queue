package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	priorityQueue := PriorityQueue{}
	priorityQueue.Enqueue(1, 0)
	priorityQueue.Enqueue(2, 1)
	iface, err := priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("can't dequeue when more than 1 element")
	}
	value, ok := iface.(int)
	if !ok {
		t.Errorf("not int in iface")
	}
	if value != 2 {
		t.Errorf("dequeue works wrong: want %v, got: %v", 2, value)
	}

}
