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
			t.Errorf("curr.Value() = %v, want %v", curr.Value(), item)
		}
		if i == 0 {
			if curr.Prev() != nil {
				t.Errorf("first.Prev() should return nil, got %p", curr.Prev())
			}
		} else {
			if curr.Prev().Value() != items[i-1] {
				t.Errorf("list.Prev().Value() = %v, want %v", curr.Prev().Value(), items[i-1])
			}
		}

		if i == len(items) - 1 {
			if curr.Next() != nil {
				t.Errorf("last.Next() should return nil, got %p", curr.Next())
			}
		} else {
			if curr.Next().Value() != items[i+1] {
				t.Errorf("list.Next().Value() = %v, want %v", curr.Next().Value(), items[i+1])
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
}