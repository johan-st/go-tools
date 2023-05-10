package linkedlist

import (
	"testing"
)

const (
	val1 = "stringValue here"
	val2 = 42
	val3 = 15.51
	val4 = "four"
	val5 = 5
	val6 = 0x6
	val7 = uint64(7)
)

func Test_list_Add(t *testing.T) {
	list := DoublyLinked{}
	checkHeadAndTail(t, list, nil, nil)

	list.Add(val1)
	checkHeadAndTail(t, list, val1, val1)

	list.Add(val2)
	checkHeadAndTail(t, list, val2, val1)

	list.Add(val3)
	checkHeadAndTail(t, list, val3, val1)
}

func Test_list_GetHead(t *testing.T) {
	list := DoublyLinked{}

	val := list.Head()
	comp(t, "head empty list", val, nil)

	list.Add(val1)
	list.Add(val2)
	list.Add(val3)
	val = list.Head()
	comp(t, "val", val, val3)

	val = list.Head()
	comp(t, "val", val, val2)

	val = list.Head()
	comp(t, "val", val, val1)

	val = list.Head()
	comp(t, "val", val, nil)

}

func Test_list_GetTail(t *testing.T) {
	list := DoublyLinked{}

	val := list.GetTail()
	comp(t, "head empty list", val, nil)

	list.Add(val1)
	list.Add(val2)
	list.Add(val3)
	val = list.GetTail()
	comp(t, "val", val, val1)

	val = list.GetTail()
	comp(t, "val", val, val2)

	val = list.GetTail()
	comp(t, "val", val, val3)

	val = list.GetTail()
	comp(t, "val", val, nil)
}

func Test_list_GetIndex(t *testing.T) {
	list := DoublyLinked{}
	list.Add(val7)
	list.Add(val6)
	list.Add(val5)
	list.Add(val4)
	list.Add(val3)
	list.Add(val2)
	list.Add(val1)

	val := list.GetIndex(4)
	comp(t, "index 4", val, val5)
	val = list.GetIndex(4)
	comp(t, "index 4 again", val, val6)
	val = list.GetIndex(5)
	comp(t, "index 6 (out of bonds)", val, nil)
	val = list.GetIndex(0)
	comp(t, "index 6 (out of bonds)", val, val1)
}

// HELPERS
func checkHeadAndTail(t *testing.T, l DoublyLinked, headWanted any, tailWanted any) {
	t.Helper()
	fail := false
	if l.PeekHead() != headWanted {
		t.Logf("head was %v, wanted %v", l.PeekHead(), headWanted)
		fail = true
	}
	if l.PeekTail() != tailWanted {
		t.Logf("tail was %v, wanted %v", l.PeekTail(), tailWanted)
		fail = true
	}
	if fail {
		t.Fail()
	}
}
func comp(t *testing.T, valName string, valWas any, valWant any) {
	t.Helper()
	if valWas != valWant {
		t.Logf("%s was %v, wanted %v", valName, valWas, valWant)
		t.Fail()
	}
}
