package priority_queue

import (
	"fmt"
	"sort"
)

type PriorityQueue struct {
	list List
}

func (queue *PriorityQueue) Len() int {
	return queue.list.Len()
}

func (queue *PriorityQueue) First() *Node {
	return queue.list.FirstPtr()
}

func (queue *PriorityQueue) Last() *Node {
	return queue.list.LastPtr()
}

func (queue *PriorityQueue) Enqueue(element interface{}, priority int) {
	queue.list.Add(element, priority)
}

func (queue *PriorityQueue) Dequeue() (interface{}, error) {
	sort.Stable(&queue.list)
	value, err := queue.list.LastValue()
	if err != nil {
		return nil, err
	}
	queue.list.DeleteLast()
	return value, nil
}

// ---------------List------------------

type Node struct {
	prev     *Node
	next     *Node
	value    interface{}
	priority int
	index    int
}

type List struct {
	firstPtr *Node
	length   int
}

func (list *List) Add(value interface{}, priority int) {
	defer func(){list.length++}()
	if list.firstPtr == nil {
		list.firstPtr = &Node{
			value:    value,
			priority: priority,
			index:    list.length,
		}
		list.firstPtr.prev = list.firstPtr
		list.firstPtr.next = list.firstPtr
		return
	}
	secondPtr := list.firstPtr
	list.firstPtr = &Node{
		value:    value,
		priority: priority,
		index:    list.length,
		prev:     secondPtr.prev,
		next:     secondPtr,
	}
	secondPtr.prev = list.firstPtr

	lastPtr := list.firstPtr.prev
	lastPtr.next = list.firstPtr
}

func (list *List) LastValue() (interface{}, error) {
	if list.firstPtr == nil {
		return nil, fmt.Errorf("no elements in list")
	}
	return list.firstPtr.prev.value, nil
}

func (list *List) DeleteLast() {
	if list.firstPtr == nil {
		return
	}
	defer func() { list.length-- }()
	if list.length == 1 {
		list.firstPtr = nil
		return
	}
	lastPtr := list.firstPtr.prev
	beforeLastPtr := lastPtr.prev
	beforeLastPtr.next = list.firstPtr
	list.firstPtr.prev = beforeLastPtr
}

func (list *List) FirstPtr() *Node {
	return list.firstPtr
}

func (list *List) LastPtr() *Node {
	return list.firstPtr.prev
}

func (list *List) Len() int {
	return list.length
}

func (list *List) Less(i, j int) bool {
	iElemPtr := helpSearchElemByIndex(list, i)
	jElemPtr := helpSearchElemByIndex(list, j)
	return iElemPtr.priority > jElemPtr.priority
}

func (list *List) Swap(i, j int) {
	iElemPtr := helpSearchElemByIndex(list, i)
	jElemPtr := helpSearchElemByIndex(list, j)
	iElemPtr.value, jElemPtr.value = jElemPtr.value, iElemPtr.value
}

func helpSearchElemByIndex(list *List, i int) *Node {
	if i >= list.length || i < 0 {
		panic("index is wrong")
	}
	iSearchPtr := list.firstPtr
	for iSearchPtr.index != i {
		iSearchPtr = iSearchPtr.next
	}
	return iSearchPtr
}