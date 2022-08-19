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
	frontEl *ListItem
	backEl  *ListItem
	len     int
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.frontEl
}

func (list *list) Back() *ListItem {
	return list.backEl
}

func (list *list) PushFront(v interface{}) *ListItem {
	frontEl := &ListItem{v, list.frontEl, nil}
	if list.Len() == 0 {
		list.backEl = frontEl
	} else {
		frontEl.Next.Prev = frontEl
	}
	list.frontEl = frontEl
	list.len++

	return list.frontEl
}

func (list *list) PushBack(v interface{}) *ListItem {
	backEl := &ListItem{v, nil, list.backEl}
	if list.Len() == 0 {
		list.frontEl = backEl
	} else {
		backEl.Prev.Next = backEl
	}
	list.backEl = backEl
	list.len++

	return list.backEl
}

func (list *list) Remove(i *ListItem) {
	if i == list.Front() {
		list.frontEl = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i == list.Back() {
		list.backEl = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	list.len--
}

func (list *list) MoveToFront(i *ListItem) {
	if list.Front() == i {
		return
	}

	if list.Back() == i {
		list.backEl = i.Prev
		i.Prev.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	list.frontEl.Prev = i
	i.Prev = nil
	i.Next = list.frontEl
	list.frontEl = i
}

func NewList() List {
	return new(list)
}
