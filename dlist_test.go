package dlist

import "testing"

func checkIntList(t *testing.T, l *List, items []int) {
	if l.Len() != len(items) {
		t.Fatalf("list.Len() = %v, want %v", l.Len(), len(items))
	}

	if len(items) == 0 {
		if l.First() != nil || l.Last() != nil {
			t.Fatalf("Last() = %v and First() = %v; both should return nil for empty list", l.Last(), l.First())
		}
		return
	}

	curr := l.First()

	for i, item := range items {
		if curr.Value() != item {
			t.Fatalf("curr.Value() = %v, want %v", curr.Value(), item)
		}
		if i == 0 {
			if curr.Prev() != nil {
				t.Fatalf("first.Prev() should return nil, got %p", curr.Prev())
			}
			if l.First() != curr {
				t.Fatalf("First() = %p, want %p", l.First(), curr)
			}
		} else {
			if curr.Prev().Value() != items[i-1] {
				t.Fatalf("list.Prev().Value() = %v, want %v", curr.Prev().Value(), items[i-1])
			}
		}

		if i == len(items) - 1 {
			if curr.Next() != nil {
				t.Fatalf("last.Next() should return nil, got %p", curr.Next())
			}
			if l.Last() != curr {
				t.Fatalf("Last() = %p, want %p", l.Last(), curr)
			}
		} else {
			if curr.Next().Value() != items[i+1] {
				t.Fatalf("list.Next().Value() = %v, want %v", curr.Next().Value(), items[i+1])
			}
		}
		curr = curr.Next()
	}
}

func TestList(t *testing.T) {
	list := &List{}
	checkIntList(t, list, []int{})
	list.PushFront(1)
	checkIntList(t, list, []int{1})
	list.PushFront(3)
	list.PushFront(5)
	checkIntList(t, list, []int{5, 3, 1})
	
	list = &List{}
	list.PushBack(3)
	checkIntList(t, list, []int{3})
	list.PushBack(5)
	list.PushBack(8)
	checkIntList(t, list, []int{3, 5, 8})
}

func TestItemRemove(t *testing.T) {
	list := &List{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)

	list.First().Remove()
	checkIntList(t, list, []int{2, 3, 4, 5, 6})
	list.Last().Remove()
	checkIntList(t, list, []int{2, 3, 4, 5})
	list.First().Next().Remove()
	checkIntList(t, list, []int{2, 4, 5})
}