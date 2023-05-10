package linkedlist_test

import (
	"testing"
)

func Test_list(t *testing.T) {
	// arrange
	val1 := "string value"
	val2 := 5
	// val3 := struct {
	// 	str     string
	// 	integer int
	// }{"bob", 001}

	// EMPTY LIST
	list := DoublyLinkedList{}
	if list.tail != nil && list.head != nil {
		t.Log("empty list should not have a head or tail")
		t.Logf("head: %#+v, tail %#+v:, want: %#+v", list.head, list.tail, nil)
		t.Fail()
	}

	// LIST WITH ONE VALUE
	list.Add(val1)
	if list.tail == nil && list.head == nil {
		t.Log("non-empty list should have a head and a tail")
		t.Logf("head: %#+v, tail %#+v:, want: %#+v", list.head, list.tail, nil)
		t.FailNow()
	}

	if list.tail.value != val1 && list.head.value != val1 {
		t.Log("list of one item should have that node as head and tail")
		t.Logf("head: %#+v, tail %#+v:, want: %#+v", list.head.value, list.tail.value, val1)
		t.Fail()
	}

	// LIST WITH TWO VALUES
	list.Add(val2)
	if list.tail.value != val1 && list.head.value != val1 {
		t.Log("tail should be the value added first and head the most recently added value")
		t.Logf("head: %#+v, want %#+v:\ntail:%#+v,  want: %#+v", list.head.value, val2, list.tail.value, val1)
		t.Fail()
	}

	if list.PeekHead() != val2 {
		t.Fatalf()
	}
}
