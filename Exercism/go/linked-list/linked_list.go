package linkedlist

import "errors"

var ErrEmptyList error = errors.New("")

// Define List and Node types here.
type List struct {
    first *Node
    last *Node
    size int
}

type Node struct {
    Val interface{}
    next *Node
    prev *Node
}

func NewList(args ...interface{}) *List {
	list := &List{}
    for _,v := range args {
        list.PushBack(v)
    }
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	node := Node{Val:v, next:l.first}
    if l.size == 0 {
        l.first, l.last = &node, &node
    } else {
    	l.first.prev = &node
        l.first = &node
    }
	l.size++
}

func (l *List) PushBack(v interface{}) {
	node := Node{Val:v, prev:l.last}
    if l.size == 0 {
        l.first, l.last = &node, &node
    } else {
    	l.last.next = &node
        l.last = &node
    }
	l.size++
}

func (l *List) PopFront() (interface{}, error) {
	if l.size == 0 {
        return nil, ErrEmptyList
    }
	if l.size == 1 {
        v := l.first.Val
        l.first, l.last = nil, nil
        l.size--
        return v, nil
    }
	v := l.first.Val
    l.first = l.first.next
    l.first.prev = nil
    l.size--
    return v, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.size == 0 {
        return nil, ErrEmptyList
    }
	if l.size == 1 {
        v := l.last.Val
        l.first, l.last = nil, nil
        l.size--
        return v, nil
    }
	v := l.last.Val
    l.last = l.last.prev
    l.last.next = nil
    l.size--
    return v, nil
}

func (l *List) Reverse() {
	data := []interface{}{}
	d, err := l.PopBack()
	for err == nil {
		data = append(data, d)
		d, err = l.PopBack()
	}
	newList := NewList(data...)
	l.first = newList.first
	l.last = newList.last
	l.size = newList.size
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
