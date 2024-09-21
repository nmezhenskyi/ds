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

package stack

import "testing"

func TestStackNew(t *testing.T) {
	st := New[int]()

	if st.size != 0 {
		t.Errorf("Expected st.size to be %d, got %d instead", 0, st.size)
	}
	if st.stack == nil {
		t.Fatal("Expected st.stack to be initialized, got nil instead")
	}
}

func TestStackPush(t *testing.T) {
	st := New[int]()

	st.Push(10)
	st.Push(20)
	st.Push(30)

	expected := []int{10, 20, 30}

	if st.Size() != 3 {
		t.Errorf("Expected st.Size() to be %d, got %d instead", 3, st.Size())
	}

	for i, v := range st.stack {
		if v != expected[i] {
			t.Fatalf("Expected element at idx %d to be %d, got %d instead",
				i, expected[i], v)
		}
	}
}

func TestStackPop(t *testing.T) {
	st := New[int]()

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Pop()

	if st.Size() != 2 {
		t.Errorf("Expected st.Size() to be %d, got %d instead", 2, st.Size())
	}
	if v := st.Top(); v != 20 {
		t.Fatalf("Expected st.Top() to return %d, got %d instead", 20, st.Top())
	}
}

func TestStackTop(t *testing.T) {
	st := New[int]()

	st.Push(10)
	if v := st.Top(); v != 10 {
		t.Fatalf("Expected st.Top() to return %d, got %d instead", 10, st.Top())
	}

	st.Push(20)
	if v := st.Top(); v != 20 {
		t.Fatalf("Expected st.Top() to return %d, got %d instead", 20, st.Top())
	}
}

func TestStackSize(t *testing.T) {
	st := New[int]()

	if st.Size() != 0 {
		t.Fatalf("Expected st.Size() to be %d, got %d instead", 0, st.Size())
	}

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(40)
	st.Push(50)

	if st.Size() != 5 {
		t.Fatalf("Expected st.Size() to be %d, got %d instead", 5, st.Size())
	}

	st.Pop()
	st.Pop()

	if st.Size() != 3 {
		t.Fatalf("Expected st.Size() to be %d, got %d instead", 3, st.Size())
	}
}

func TestStackIsEmpty(t *testing.T) {
	st := New[int]()

	if !st.IsEmpty() {
		t.Fatal("Expected st to be empty")
	}

	st.Push(1)

	if st.IsEmpty() {
		t.Fatal("Expected st to be non-empty")
	}
}
