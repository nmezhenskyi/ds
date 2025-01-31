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

	l.Insert(0, 10)
	l.Insert(0, 20)
	l.Insert(1, 30)

	if n := l.Size(); n != 3 {
		t.Fatalf("Expected size to be %d, got %d instead\n", 3, n)
	}

	if v, _ := l.Find(0); v != 20 {
		t.Errorf("Expected l[0] to be %d, got %d instead\n", 20, v)
	}
	if v, _ := l.Find(1); v != 10 {
		t.Errorf("Expected l[1] to be %d, got %d instead\n", 10, v)
	}
	if v, _ := l.Find(2); v != 30 {
		t.Errorf("Expected l[2] to be %d, got %d instead\n", 30, v)
	}
}

func TestListInsertAtNegativeIdx(t *testing.T) {
	l := List[int]{}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected to recover from panic after l.Insert(-1, 10), got nil instead")
		}
	}()

	l.Insert(-1, 10)
}

func TestInsertAtHighIdx(t *testing.T) {
	l := List[int]{}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected to recover from panic after l.Insert(-1, 10), got nil instead")
		}
	}()

	l.Insert(100, 10)
}

func TestListReplace(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	l.Replace(1, 1000)

	if _, ok := l.Find(1); !ok {
		t.Errorf("Expected ok to be true, got %t instead\n", ok)
	}
	if v, _ := l.Find(1); v != 1000 {
		t.Fatalf("Expected l.Find(1) to return %d, got %d instead\n", 1000, v)
	}
}

func TestListFind(t *testing.T) {
	l := List[int]{}

	l.Append(10)

	v, ok := l.Find(0)
	if !ok {
		t.Errorf("Expected l.Find(0) ok to be true after l.Append(10), got %t instead\n", ok)
	}
	if v != 10 {
		t.Errorf("Expected l.Find(0) v to be %d after l.Append(10), got %d instead\n", 10, v)
	}

	l.Append(20)
	l.Append(30)
	l.Append(40)

	v, ok = l.Find(0)
	if !ok {
		t.Errorf("Expected l.Find(0) ok to be true, got %t instead\n", ok)
	}
	if v != 10 {
		t.Errorf("Expected l.Find(0) v to be %d, got %d instead\n", 10, v)
	}

	v, ok = l.Find(1)
	if !ok {
		t.Errorf("Expected l.Find(0) ok to be true, got %t instead\n", ok)
	}
	if v != 20 {
		t.Errorf("Expected l.Find(0) v to be %d, got %d instead\n", 20, v)
	}
}

func TestListIndexOf(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	if i := l.IndexOf(10); i != 0 {
		t.Errorf("Expected l.IndexOf(10) to return %d, got %d instead\n", 0, l.IndexOf(10))
	}
	if i := l.IndexOf(20); i != 1 {
		t.Errorf("Expected l.IndexOf(20) to return %d, got %d instead\n", 1, l.IndexOf(20))
	}
	if i := l.IndexOf(30); i != 2 {
		t.Errorf("Expected l.IndexOf(30) to return %d, got %d instead\n", 2, l.IndexOf(30))
	}
	if i := l.IndexOf(40); i != -1 {
		t.Errorf("Expected l.IndexOf(40) to return %d, got %d instead\n", -1, l.IndexOf(40))
	}
}

func TestListContains(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	if ok := l.Contains(10); !ok {
		t.Errorf("Expected l.Contains(10) to return true, got %t instead\n", ok)
	}
	if ok := l.Contains(20); !ok {
		t.Errorf("Expected l.Contains(20) to return true, got %t instead\n", ok)
	}
	if ok := l.Contains(30); !ok {
		t.Errorf("Expected l.Contains(30) to return true, got %t instead\n", ok)
	}
	if ok := l.Contains(40); ok {
		t.Errorf("Expected l.Contains(40) to return false, got %t instead\n", ok)
	}
}

func TestListRemove(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	l.Remove(1)

	if _, ok := l.Find(1); ok {
		t.Errorf("Expected l.Find(1) ok to be false after l.Remove(), got %t instead\n", ok)
	}
	if l.Size() != 2 {
		t.Errorf("Expected l.Size() to be %d, got %d instead\n", 2, l.Size())
	}
}

func TestListClear(t *testing.T) {
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)

	l.Clear()

	if n := l.Size(); n != 0 {
		t.Fatalf("Expected size to be %d, got %d instead\n", 0, n)
	}
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

	expected := []int{100, 20, 30, 40, 50, 60, 70, 80, 90, 10}
	res := l.ToSlice()

	if size, expSize := l.Size(), len(expected); size != expSize {
		t.Fatalf("Expected list size to be %d, got %d instead\n", len(expected), l.Size())
	}

	for i, n := 0, l.Size(); i < n; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected res[%d] to be %d, got %d instead\n", i, expected[i], res[i])
		}
	}
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
	l := List[int]{}

	l.Append(10)
	l.Append(20)
	l.Append(30)
	l.Append(40)
	l.Append(50)

	expected := []int{10, 20, 30, 40, 50}
	res := l.ToSlice()

	if size, expSize := l.Size(), len(expected); size != expSize {
		t.Fatalf("Expected list size to be %d, got %d instead\n", len(expected), l.Size())
	}

	for i, n := 0, l.Size(); i < n; i++ {
		if res[i] != expected[i] {
			t.Errorf("Expected res[%d] to be %d, got %d instead\n", i, expected[i], res[i])
		}
	}
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
