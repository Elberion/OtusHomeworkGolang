package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head    *ListItem
	tail    *ListItem
	listLen int
}

func (l *list) Len() int {
	return l.listLen
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) pushFirst(v interface{}) *ListItem {
	CurrentNode := ListItem{v, nil, nil}
	l.head = &CurrentNode
	l.tail = &CurrentNode
	l.listLen = 1
	return &CurrentNode
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.listLen == 0 {
		return l.pushFirst(v)
	}
	CurrentNode := ListItem{v, l.head, nil}
	l.head.Prev = &CurrentNode
	l.head = &CurrentNode
	l.listLen++
	return &CurrentNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.listLen == 0 {
		return l.pushFirst(v)
	}
	CurrentNode := ListItem{v, nil, l.tail}
	l.tail.Next = &CurrentNode
	l.tail = &CurrentNode
	l.listLen++
	return &CurrentNode
}

func (l *list) Remove(i *ListItem) {
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
		l.listLen = 0
		return
	}
	l.listLen--
	if i == l.head {
		l.head = i.Next
		l.head.Prev = nil
		return
	}
	if i == l.tail {
		l.tail = i.Prev
		l.tail.Next = nil
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.head {
		if i == l.tail {
			i.Prev.Next = nil
			l.tail = i.Prev
		} else {
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
		}
		i.Next = l.head
		i.Prev = nil
		l.head.Prev = i
		l.head = i
	}
}

func NewList() List {
	return new(list)
}
