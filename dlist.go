package dlist

// Item is an item of a linked list
type Item struct {
	value      interface{}
	next, prev *Item
	list       *List
}

// Value returns a value of item
func (it *Item) Value() interface{} { return it.value }

// Next returns next list item or nil.
func (it *Item) Next() *Item { return it.next }

// Prev returns previous list item or nil.
func (it *Item) Prev() *Item { return it.prev }

// List is a doubly linked list
type List struct {
	head, tail *Item
	count      int
}

// Len returns the number of elements of list.
func (l *List) Len() int { return l.count }

// First returns the first element or nil
func (l *List) First() *Item { return l.head }

// Last return the last element or nil
func (l *List) Last() *Item { return l.tail }

// PushFront inserts a new item with given value at the front of list.
func (l *List) PushFront(v interface{}) {
	newItem := &Item{v, nil, nil, l}
	if l.count == 0 {
		l.tail = newItem
		l.head = newItem
	} else {
		l.head.prev = newItem
		newItem.next = l.head
		l.head = newItem
	}
	l.count++
}
