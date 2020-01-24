package priority_queue

import (
	"testing"
)

func TestPriorityQueue_Empty(t *testing.T) {
	priorityQueue := PriorityQueue{}

	firstPtr := priorityQueue.First()
	if firstPtr != nil {
		t.Errorf("First() must be nil, got %v", firstPtr)
	}

	lastPtr := priorityQueue.Last()
	if lastPtr != nil {
		t.Errorf("Last() must be nil, got: %v", lastPtr)
	}

	length := priorityQueue.Len()
	if length != 0 {
		t.Errorf("Len() must be 0, got: %v", length)
	}

	_, err := priorityQueue.Dequeue()
	if err == nil {
		t.Errorf("method Dequeue() must return not nil error")
	}
}

func TestPriorityQueue_OneElement(t *testing.T) {
	priorityQueue := PriorityQueue{}
	priorityQueue.Enqueue(123, 0)

	length := priorityQueue.Len()
	if length != 1 {
		t.Errorf("length must be 1, got: %v", length)
	}

	firstPtr := priorityQueue.First()
	if firstPtr != priorityQueue.lst.firstPtr {
		t.Errorf("First() must be equal to the added element")
	}

	lastPtr := priorityQueue.Last()
	if lastPtr != priorityQueue.lst.firstPtr.prev {
		t.Errorf("Last() must be equal to the added element")
	}

	valueFace, err := priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("err must be nil, got: %v", err)
	}

	value, ok := valueFace.(int)
	if !ok {
		t.Fatal("valueFace must be int")
	}
	if value != 123 {
		t.Errorf("value must be 123, got: %v", value)
	}

	length = priorityQueue.Len()
	if length != 0 {
		t.Errorf("length must be 0, got: %v", length)
	}
}

func TestPriorityQueue_ManyElements_NoPriority(t *testing.T) {
	priorityQueue := PriorityQueue{}
	priorityQueue.Enqueue(123, 0)
	priorityQueue.Enqueue(456, 0)

	length := priorityQueue.Len()
	if length != 2 {
		t.Errorf("length must be 2, got: %v", length)
	}

	firstPtr := priorityQueue.First()
	if firstPtr != priorityQueue.lst.firstPtr {
		t.Errorf("First() must be equal to the added element")
	}

	lastPtr := priorityQueue.Last()
	if lastPtr != priorityQueue.lst.firstPtr.prev {
		t.Errorf("Last() must be equal to the added element")
	}

	valueFace, err := priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("err must be nil, got: %v", err)
	}

	value, ok := valueFace.(int)
	if !ok {
		t.Fatal("valueFace must be type int")
	}

	if value != 123 {
		t.Errorf("value must be 123, got: %v", value)
	}

	valueFace, err = priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("err must be nil, got: %v", err)
	}

	value = valueFace.(int)

	if value != 456 {
		t.Errorf("value must be 456, got: %v", value)
	}

	length = priorityQueue.Len()
	if length != 0 {
		t.Errorf("length must be 0, got: %v", length)
	}
}

func TestPriorityQueue_ManyElements_Priority(t *testing.T) {
	priorityQueue := PriorityQueue{}
	priorityQueue.Enqueue(123, 0)
	priorityQueue.Enqueue(456, 1)

	length := priorityQueue.Len()
	if length != 2 {
		t.Errorf("length must be 2, got: %v", length)
	}

	firstPtr := priorityQueue.First()
	if firstPtr != priorityQueue.lst.firstPtr {
		t.Errorf("First() must be equal to the added element")
	}

	lastPtr := priorityQueue.Last()
	if lastPtr != priorityQueue.lst.firstPtr.prev {
		t.Errorf("Last() must be equal to the added element")
	}

	valueFace, err := priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("err must be nil, got: %v", err)
	}

	value, ok := valueFace.(int)
	if !ok {
		t.Fatal("valueFace must be type int")
	}

	if value != 456 {
		t.Errorf("value must be 456, got: %v", value)
	}

	valueFace, err = priorityQueue.Dequeue()
	if err != nil {
		t.Errorf("err must be nil, got: %v", err)
	}

	value = valueFace.(int)

	if value != 123 {
		t.Errorf("value must be 123, got: %v", value)
	}

	length = priorityQueue.Len()
	if length != 0 {
		t.Errorf("length must be 0, got: %v", length)
	}
}

func TestPriorityQueue_IndexesChecking(t *testing.T) {
	q := PriorityQueue{}
	q.Enqueue("Alisher", 10)
	q.Enqueue("Boba", 9)
	q.Enqueue("Cily", 8)

	dequeueTest(&q, "Alisher", t)

	q.Enqueue("Dadpul", 7)

	dequeueTest(&q, "Boba", t)
	dequeueTest(&q, "Cily", t)
	dequeueTest(&q, "Dadpul", t)

	if q.Len() != 0 {
		t.Errorf("q.Len() must be 0, got: %v", q.Len())
	}
}
func TestPriorityQueue_StabilityChecking(t *testing.T) {
	q := PriorityQueue{}
	q.Enqueue("1", 1)
	q.Enqueue("2", 1)
	q.Enqueue("3", 1)

	dequeueTest(&q, "1", t)
	dequeueTest(&q, "2", t)
	dequeueTest(&q, "3", t)
}
func dequeueTest(q *PriorityQueue, want string, t *testing.T) {
	valueFace, _ := q.Dequeue()
	got := valueFace.(string)
	if got != want {
		t.Errorf("value must be \"%v\", got: %v", want, got)
	}
}
