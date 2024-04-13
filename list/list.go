/*
	The MIT License (MIT)

	Copyright (c) 2024 Nikita Mezhenskyi

	Permission is hereby granted, free of charge, to any person obtaining a copy of this software
	and associated documentation files (the "Software"), to deal in the Software without restriction,
	including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
	and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
	subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all copies or
	substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// Package list implements a singly-linked list.
package list

type List[V comparable] struct {
	head *node[V]
	tail *node[V]
	size int
}

type node[V comparable] struct {
	next  *node[V]
	value V
}

func (l *List[V]) Append(val V) {
	if l == nil {
		panic("list: called Append() on a nil list")
	}

	newNode := &node[V]{value: val}

	if l.size == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.size++
}

func (l *List[V]) Prepend(val V) {
	if l == nil {
		panic("list: called Prepend() on a nil list")
	}

	newNode := &node[V]{value: val, next: l.head}
	l.head = newNode

	if l.size == 0 {
		l.tail = newNode
	}

	l.size++
}

func (l *List[V]) Insert(idx int, val V) {
	if l == nil {
		panic("list: called Insert() on a nil list")
	}

	newNode := &node[V]{value: val}

	if l.size == 0 {
		// add as first element:
		l.head = newNode
		l.tail = newNode
	} else if idx >= l.size-1 {
		// append to the end:
		l.tail.next = newNode
		l.tail = newNode
	} else {
		// insert at index:
		i, curr := 0, l.head
		for curr != nil {
			if i == idx {
				temp := curr.next
				newNode.next = temp
				curr.next = newNode
				break
			}
			curr = curr.next
			i++
		}
	}

	l.size++
}

func (l *List[V]) Replace(idx int, val V) {
	if l == nil {
		panic("list: called Replace() on a nil list")
	}
	if l.size == 0 {
		return
	}

	i, curr := 0, l.head
	for curr != nil {
		if i == idx {
			curr.value = val
			return
		}
		curr = curr.next
		i++
	}
}

func (l *List[V]) Find(idx int) (v V, ok bool) {
	if l == nil {
		panic("list: called Find() on a nil list")
	}
	if l.size == 0 {
		return v, false
	}

	i, curr := 0, l.head
	for curr != nil {
		if i == idx {
			return curr.value, true
		}
		curr = curr.next
		i++
	}

	return v, false
}

func (l *List[V]) IndexOf(val V) int {
	if l == nil {
		panic("list: called IndexOf() on a nil list")
	}
	if l.size == 0 {
		return -1
	}

	idx, curr := 0, l.head
	for curr != nil {
		if curr.value == val {
			return idx
		}
		curr = curr.next
		idx++
	}

	return -1
}

func (l *List[V]) Contains(val V) bool {
	if l == nil {
		panic("list: calling Contains() on a nil list")
	}

	for curr := l.head; curr != nil; curr = curr.next {
		if curr.value == val {
			return true
		}
	}

	return false
}

func (l *List[V]) Remove(idx int) {
	if l == nil {
		panic("list: caling Remove() on a nil list")
	}
	if l.size == 0 {
		return
	}
	if l.size == 1 {
		if idx != 0 {
			return
		}

		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}

	i, prev, curr := 0, l.head, l.head.next
	for curr != nil {
		if i == idx {
			prev.next = curr.next
			curr = nil
			l.size--
			return
		}
		curr = curr.next
		i++
	}
}

func (l *List[V]) Clear() {
	if l == nil {
		return
	}
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List[V]) Swap(idx1, idx2 int) {
	if l == nil {
		panic("list: calling Swap() on a nil list")
	}

	if idx1 >= l.size || idx1 < 0 {
		return
	}
	if idx2 >= l.size || idx2 < 0 {
		return
	}
	if idx1 == idx2 {
		return
	}

	var prev1, node1 *node[V] = nil, l.head
	i1 := 0
	for node1 != nil {
		if i1 == idx1 {
			break
		}
		prev1 = node1
		node1 = node1.next
		i1++
	}

	var prev2, node2 *node[V] = nil, l.head
	i2 := 0
	for node2 != nil {
		if i2 == idx2 {
			break
		}
		prev2 = node2
		node2 = node2.next
		i2++
	}

	if prev1 != nil {
		prev1.next = node2
	}
	temp := node2.next
	node2.next = node1.next

	if prev2 != nil {
		prev2.next = node1
	}
	node1.next = temp

	if i1 == 0 {
		l.head = node2
	} else if i2 == 0 {
		l.head = node1
	}
}

func (l *List[V]) ToSlice() []V {
	if l == nil {
		panic("list: calling ToSlice() on a nil list")
	}
	if l.size == 0 {
		return nil
	}

	out := make([]V, l.size)
	i := 0
	for curr := l.head; curr != nil; curr = curr.next {
		out[i] = curr.value
		i++
	}

	return out
}

func (l *List[V]) Size() int {
	if l == nil {
		return 0
	}
	return l.size
}

func (l *List[V]) Empty() bool {
	return l == nil || l.size == 0
}
