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

package list

import "testing"

func TestListAppend(t *testing.T) {
	l := List[int]{}

	l.Append(14)

	if n := l.Size(); n != 1 {
		t.Fatalf("Expected list size to be 1, got %d instead\n", n)
	}
	if l.head == nil {
		t.Fatal("List head is nil")
	}
	if l.tail == nil {
		t.Fatal("List tail is nil")
	}
	if v, ok := l.Find(0); !ok || v != 14 {
		t.Fatalf("Expected value to be 14, got %d instead\n", v)
	}

	l.Append(15)

	if n := l.Size(); n != 2 {
		t.Fatalf("Expected list size to be 2, got %d instead\n", n)
	}
	if v, ok := l.Find(1); !ok || v != 15 {
		t.Fatalf("Expected value to be 15, got %d instead\n", v)
	}
}

func TestListPrepend(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Prepend(15)

	if n := l.Size(); n != 2 {
		t.Fatalf("Expected list size to be 2, got %d instead\n", n)
	}
	if v, ok := l.Find(0); !ok || v != 15 {
		t.Fatalf("Expected value to be 15, got %d instead\n", v)
	}
}

func TestListInsert(t *testing.T) {
	l := List[int]{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected code to panic")
		}
	}()

	l.Insert(-1, 5)

	if n := l.Size(); n != 0 {
		t.Errorf("Expected Insert(-1,5) to result in size %d, got %d instead\n",
			0, l.Size())
	}

	// TODO
}

func TestListReplace(t *testing.T) {

}

func TestListFind(t *testing.T) {

}

func TestListIndexOf(t *testing.T) {

}

func TestListContains(t *testing.T) {

}

func TestListRemove(t *testing.T) {

}

func TestListClear(t *testing.T) {

}

func TestListSwap(t *testing.T) {
	l := List[int]{}

	l.Append(10)  // 0
	l.Append(20)  // 1
	l.Append(30)  // 2
	l.Append(40)  // 3
	l.Append(50)  // 4
	l.Append(60)  // 5
	l.Append(70)  // 6
	l.Append(80)  // 7
	l.Append(90)  // 8
	l.Append(100) // 9

	l.Swap(0, 9)

	// TODO: continue
}

func TestListReverse(t *testing.T) {
	l := List[int]{}

	l.Append(10)  // 0
	l.Append(20)  // 1
	l.Append(30)  // 2
	l.Append(40)  // 3
	l.Append(50)  // 4
	l.Append(60)  // 5
	l.Append(70)  // 6
	l.Append(80)  // 7
	l.Append(90)  // 8
	l.Append(100) // 9

	l.Reverse()

	expected := []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}
	res := l.ToSlice()

	if size, expSize := l.Size(), len(expected); size != expSize {
		t.Fatalf("Expected list size to be %d, got %d instead\n", len(expected), l.Size())
	}

	for i, n := 0, l.Size(); i < n; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected res[%d] to be %d, got %d instead\n", i, expected[i], res[i])
		}
	}

	if l.head == nil {
		t.Fatal("List head is nil")
	}
	if l.head.value != 100 {
		t.Errorf("Expected list head to be 100, got %d instead\n", l.head.value)
	}
	if l.tail == nil {
		t.Fatal("List tail is nil")
	}
	if l.tail.value != 10 {
		t.Errorf("Expected list tail to be 10, got %d instead\n", l.tail.value)
	}
}

func TestListToSlice(t *testing.T) {

}

func TestListSize(t *testing.T) {
	l := List[int]{}

	if l.Size() != 0 {
		t.Errorf("Expected l.Size() on empty list to be %d, got %d instead\n", 0, l.Size())
	}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	if l.Size() != 3 {
		t.Errorf("Expected l.Size() after appends to be %d, got %d instead\n", 3, l.Size())
	}

	l.Remove(0)

	if l.Size() != 2 {
		t.Errorf("Expected l.Size() after remove to be %d, got %d instead\n", 2, l.Size())
	}
}

func TestListIsEmpty(t *testing.T) {
	l := List[int]{}

	if !l.IsEmpty() {
		t.Fatalf("Expected new list to be empty, got %t instead\n", l.IsEmpty())
	}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	if l.IsEmpty() {
		t.Fatalf("Expected list to be non-empty, got %t instead\n", l.IsEmpty())
	}
}
