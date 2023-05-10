package linkedlist

import "fmt"

type node struct {
	prev  *node
	next  *node
	value any
}

type DoublyLinked struct {
	head *node
	tail *node
	len  int
}

// Add a new node to the front of the list
func (l *DoublyLinked) Add(val any) {
	// create node
	n := node{
		prev:  nil,
		next:  l.head,
		value: val,
	}
	// set links
	if l.head != nil {
		l.head.prev = &n
	}
	l.head = &n
	if l.len == 0 {
		l.tail = l.head
	}

	// increment length
	l.len++
}

// O(1)
func (l *DoublyLinked) PeekHead() any {
	if l.head == nil {
		return nil
	}
	return l.head.value
}

// O(1)
func (l *DoublyLinked) PeekTail() any {
	if l.tail == nil {
		return nil
	}
	return l.tail.value
}

// O(1)
// Head removes the current head and returns the value contained
func (l *DoublyLinked) Head() any {
	if l.head == nil {
		return nil
	}
	val := l.head.value
	fmt.Printf("l: %v\n", l)
	fmt.Printf("h: %v\n", l.head)
	fmt.Printf("t: %v\n## detatch -----------------------------------\n", l.tail)
	l.detatch(l.head)
	fmt.Printf("l: %v\n", l)
	fmt.Printf("h: %v\n", l.head)
	fmt.Printf("t: %v\n\n", l.tail)
	return val
}

// O(1)
// Tail removes the current tail and returns the value contained
func (l *DoublyLinked) GetTail() any {
	if l.tail == nil {
		return nil
	}
	val := l.tail.value
	l.detatch(l.tail)
	return val
}

// O(n)
// GetIndex removes the given index and returns the value contained
func (l *DoublyLinked) GetIndex(index int) any {
	if index >= l.len {
		return nil
	}
	if index > l.len/2 {
		node := l.tail
		for i := l.len - 1; i >= index; i-- {
			if i == index {
				val := node.value
				l.detatch(node)
				return val
			}
			node = node.prev
		}
	} else {
		node := l.head
		for i := 0; i <= index; i++ {
			if i == index {
				val := node.value
				l.detatch(node)
				return val
			}
			node = node.next
		}
	}

	return nil
}

// O(1)
func (l *DoublyLinked) detatch(n *node) {

	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n == l.head {
		l.head = n.next
	}

	if n == l.tail {
		l.tail = n.prev
	}

	// does this do anything?
	n.next = nil
	n.prev = nil

	l.len--
}
