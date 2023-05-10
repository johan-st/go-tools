package linkedlist

type node struct {
	prev  *node
	next  *node
	value any
}

type DoublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

// Add a new node to the list
func (l *DoublyLinkedList) Add(val any) {
	// create node
	n := node{
		prev:  nil,
		next:  l.head,
		value: val,
	}
	// set head
	l.head = &n

	// if no nodes in list head and tail will both be pointng to the new node
	if l.length == 0 {
		l.tail = l.head
	}

	// increment length
	l.length++
}

func (l *DoublyLinkedList) PeekHead() any {
	if l.head == nil {
		return nil
	}
	return l.head.value
}
