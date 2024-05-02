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

// Package rbuf implements a generic Ring Buffer.
package rbuf

type RingBuffer[T any] struct {
	buf   []T
	size  int
	count int
	read  int
	write int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	if size <= 0 {
		panic("rbuf: buffer size must be positive value")
	}

	return &RingBuffer[T]{
		buf:  make([]T, size),
		size: size,
	}
}

func (rb *RingBuffer[T]) Enqueue(val T) bool {
	if rb.count == rb.size {
		return false
	}

	rb.buf[rb.write] = val
	rb.write = (rb.write + 1) % rb.size
	rb.count++

	return true
}

func (rb *RingBuffer[T]) EnqueueOverwrite(val T) {
	// TODO: revisit this edge case.
	if rb.read == rb.write {
		rb.read = (rb.read + 1) % rb.size
	}

	rb.buf[rb.write] = val
	rb.write = (rb.write + 1) % rb.size

	if rb.count < rb.size {
		rb.count++
	}
}

func (rb *RingBuffer[T]) Dequeue() (val T, ok bool) {
	if rb.count == 0 {
		return val, false
	}

	val = rb.buf[rb.read]
	rb.read = (rb.read + 1) % rb.size
	rb.count--

	return val, true
}

func (rb *RingBuffer[T]) Peek() (val T, ok bool) {
	if rb.count == 0 {
		return val, false
	}
	return rb.buf[rb.read], true
}

func (rb *RingBuffer[T]) Emit() []T {
	if rb.count == 0 {
		return nil
	}

	out := make([]T, rb.count)

	for rb.count > 0 {
		out = append(out, rb.buf[rb.read])
		rb.read = (rb.read + 1) % rb.size
		rb.count--
	}

	return out
}

func (rb *RingBuffer[T]) IsFull() bool {
	return rb.count == rb.size
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.count == 0
}
