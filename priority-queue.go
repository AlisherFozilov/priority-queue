package priority_queue

import (
	"fmt"
	"log"
	"sort"
)

type PriorityQueue struct {
	lst list
}

func (queue *PriorityQueue) Len() int {
	return queue.lst.Len()
}

func (queue *PriorityQueue) Enqueue(element interface{}, priority int) {
	queue.lst.add(element, priority)
}

func (queue *PriorityQueue) Dequeue() (interface{}, error) {
	sort.Stable(&queue.lst)
	value, err := queue.lst.lastValue()
	if err != nil {
		return nil, err
	}
	queue.lst.deleteLast()
	return value, nil
}

func (queue *PriorityQueue) First() interface{} {
	firstPtr := queue.lst.getFirstPtr()
	if firstPtr != nil {
		return firstPtr.value
	}
	return nil
}

func (queue *PriorityQueue) Last() interface{} {
	lastPtr := queue.lst.getLastPtr()
	if lastPtr != nil {
		return lastPtr.value
	}
	return nil
}

// ---------------list------------------

type node struct {
	prev     *node
	next     *node
	value    interface{}
	priority int
	index    int
}

type list struct {
	firstPtr *node
	length   int
}

func (list *list) add(value interface{}, priority int) {
	defer func() { list.length++ }()
	if list.firstPtr == nil {
		list.firstPtr = &node{
			value:    value,
			priority: priority,
			index:    list.length,
		}
		list.firstPtr.prev = list.firstPtr
		list.firstPtr.next = list.firstPtr
		return
	}
	secondPtr := list.firstPtr
	list.firstPtr = &node{
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

func (list *list) lastValue() (interface{}, error) {
	if list.firstPtr == nil {
		return nil, fmt.Errorf("no elements in lst")
	}
	return list.firstPtr.prev.value, nil
}

func (list *list) deleteLast() {
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
	list.refreshIndexes()
}

func (list *list) getFirstPtr() *node {
	return list.firstPtr
}

func (list *list) getLastPtr() *node {
	if list.firstPtr == nil {
		return nil
	}
	return list.firstPtr.prev
}

func (list *list) Len() int {
	return list.length
}

func (list *list) Less(i, j int) bool {
	iElemPtr := list.helpSearchElemByIndex(i)
	jElemPtr := list.helpSearchElemByIndex(j)
	return iElemPtr.priority > jElemPtr.priority
}

func (list *list) Swap(i, j int) {
	iElemPtr := list.helpSearchElemByIndex(i)
	jElemPtr := list.helpSearchElemByIndex(j)
	iElemPtr.value, jElemPtr.value = jElemPtr.value, iElemPtr.value
}

func (list *list) helpSearchElemByIndex(i int) *node {
	if i >= list.length || i < 0 {
		panic("index is wrong")
	}
	iSearchPtr := list.firstPtr

	counter := 1
	for iSearchPtr.index != i {
		iSearchPtr = iSearchPtr.next
		if counter > list.length {
			log.Fatal("...")
		}
		counter++
	}
	return iSearchPtr
}

func (list *list) refreshIndexes() {
	currentPtr := list.firstPtr
	for i := 0; i < list.length-1; i++ {
		currentPtr.index--
		currentPtr = currentPtr.next
	}
}
