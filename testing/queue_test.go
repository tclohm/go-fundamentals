package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0 ; i < 3 ; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("failed to append item %v", i)
		}
	}

	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	} 
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0 ; i < 3 ; i++ {
		q.Append(i)
	}

	for i := 0 ; i < 3 ; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should not be able get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: %v vs %v", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("should not have any more items in queue: %v", item)
	}
}