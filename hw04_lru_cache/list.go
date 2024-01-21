package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

// Define the list structure
type list struct {
	// Place your code here.
	len  int
	head *ListItem
	tail *ListItem
}

// Add a new node at the end of the doubly linked list
func (dll *list) PushBack(data interface{}) *ListItem {
	newItem := &ListItem{Value: data, Prev: nil, Next: nil}

	if dll.head == nil {
		dll.head = newItem
		dll.tail = newItem
	} else {
		newItem.Prev = dll.tail
		dll.tail.Next = newItem
		dll.tail = newItem
	}

	dll.len++

	return newItem
}

// Add a new node at the beginning of the doubly linked list
func (dll *list) PushFront(data interface{}) *ListItem {
	newItem := &ListItem{Value: data, Prev: nil, Next: nil}

	if dll.head == nil {
		dll.head = newItem
		dll.tail = newItem
	} else {
		newItem.Next = dll.head
		dll.head.Prev = newItem
		dll.head = newItem
	}

	dll.len++
	return newItem
}

// Remove element i from the doubly linked list
func (dll *list) Remove(i *ListItem) {
	if i == dll.head {
		dll.head = i.Next
	}
	if i == dll.tail {
		dll.tail = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i = nil

	dll.len--
}

// Move element i to the front
func (dll *list) MoveToFront(i *ListItem) {
	dll.Remove(i)
	dll.PushFront(i.Value)
}

// Length of the doubly linked list
func (dll *list) Len() int {
	return dll.len
}

// Get the front element
func (dll *list) Front() *ListItem {
	return dll.head
}

// Get the back element
func (dll *list) Back() *ListItem {
	return dll.tail
}

// Initialize a new empty doubly linked list
func NewList() List {
	return new(list)
}
