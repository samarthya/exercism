package linkedlist

import (
	"errors"
	"log"
)

var ErrEmptyList = errors.New("no elements in the list")

// Define List and Node types here.
type Node struct {
	Val            interface{}
	next, previous *Node
}

type List struct {
	Head, Tail *Node
}

func NewList(args ...interface{}) *List {
	l := &List{}

	for _, v := range args {
		l.PushBack(v)
	}

	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) PushFront(v interface{}) {
	n := &Node{Val: v}

	if l.Head == nil {
		l.Tail = n
	} else {
		n.next = l.Head
		l.Head.previous = n
	}
	l.Head = n
}

func (l *List) PushBack(v interface{}) {
	n := &Node{Val: v}

	if l.Head == nil {
		l.Head = n
		n.previous = nil
		n.next = nil
		l.Tail = n
		return
	}

	l.Tail.next = n
	n.previous = l.Tail
	l.Tail = n
}

func (l *List) PopFront() (interface{}, error) {

	if l.Head == nil {
		return "", ErrEmptyList
	}

	m := l.Head

	if m.next == nil {
		l.Head = nil
		l.Tail = nil

	} else {
		l.Head = m.next
		l.Head.previous = nil
	}

	return m.Val, nil
}

func (l *List) PopBack() (interface{}, error) {

	if l.Tail == nil {
		return "", ErrEmptyList
	}

	m := l.Tail
	if m.previous == nil {
		l.Head = nil
		l.Tail = nil

	} else {
		l.Tail = m.previous
		l.Tail.next = nil
	}

	return m.Val, nil
}

func (l *List) Reverse() {
	c := l.Head
	for c != nil {
		c.previous, c.next = c.next, c.previous
		c = c.previous
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func (l *List) Display() {
	for m := l.First(); m.next != nil; m = m.next {
		log.Printf(" Element: %d\n", m.Val)
	}
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}
