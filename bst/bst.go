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

// Package bst implements AVL balanced binary search tree.
package bst

import (
	"cmp"
	"fmt"
)

type Tree[K cmp.Ordered, V any] struct {
	root *Node[K, V]
	size int
}

func (tree *Tree[K, V]) Insert(key K, data V) {
	if tree == nil {
		panic("bst: called Insert() on a nil tree")
	}
	tree.root = insert(tree.root, key, data, &tree.size)
}

func (tree *Tree[K, V]) Find(key K) *Node[K, V] {
	if tree == nil {
		panic("bst: called Find() on a nil tree")
	}
	return find(tree.root, key)
}

func (tree *Tree[K, V]) Remove(key K) {
	if tree == nil {
		panic("bst: called Remove() on a nil tree")
	}
	tree.root = remove(tree.root, key, &tree.size)
}

func (tree *Tree[K, V]) Height() int {
	if tree == nil || tree.root == nil {
		return 0
	}
	return tree.root.height
}

func (tree *Tree[K, V]) Size() int {
	if tree == nil {
		return 0
	}
	return tree.size
}

func (tree *Tree[K, V]) Keys() []K {
	keys := make([]K, 0)

	var walk func(node *Node[K, V])
	walk = func(node *Node[K, V]) {
		if node == nil {
			return
		}
		walk(node.left)
		keys = append(keys, node.key)
		walk(node.right)
	}
	walk(tree.root)

	return keys
}

type Node[K cmp.Ordered, V any] struct {
	key    K
	left   *Node[K, V]
	right  *Node[K, V]
	height int
	data   V
}

func (node *Node[K, V]) Key() K {
	return node.key
}

func (node *Node[K, V]) Data() V {
	return node.data
}

func (node *Node[K, V]) String() string {
	return fmt.Sprintf("%v", node.key)
}

func insert[K cmp.Ordered, V any](node *Node[K, V], key K, data V, size *int) *Node[K, V] {
	if node == nil || node.height == 0 {
		*size++
		return &Node[K, V]{
			key:    key,
			left:   nil,
			right:  nil,
			height: 1,
			data:   data,
		}
	}

	if key < node.key {
		node.left = insert(node.left, key, data, size)
	} else if key > node.key {
		node.right = insert(node.right, key, data, size)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))
	balance := getBalance(node)

	if balance > 1 && key < node.left.key {
		return rotateRight(node)
	}
	if balance < -1 && key > node.right.key {
		return rotateLeft(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}
	if balance < -1 && key < node.right.key {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	return node
}

func find[K cmp.Ordered, V any](node *Node[K, V], key K) *Node[K, V] {
	if node == nil {
		return nil
	}

	if key == node.key {
		return node
	}

	if key < node.key {
		return find(node.left, key)
	} else {
		return find(node.right, key)
	}
}

func remove[K cmp.Ordered, V any](node *Node[K, V], key K, size *int) *Node[K, V] {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key, size)
	} else if key > node.key {
		node.right = remove(node.right, key, size)
	} else {
		if node.left == nil || node.right == nil {
			var temp *Node[K, V]

			if node.left != nil {
				temp = node.left
			} else {
				temp = node.right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				// TODO: review this, does a shallow copy.
				*node = *temp
			}

			*size--
		} else {
			temp := getMinNode(node.right)
			node.key = temp.key
			node.right = remove(node.right, temp.key, size)
		}
	}

	if node == nil {
		return nil
	}

	node.height = 1 + max(height(node.left), height(node.right))
	balance := getBalance(node)

	if balance > 1 && getBalance(node.left) >= 0 {
		return rotateRight(node)
	}
	if balance > 1 && getBalance(node.left) < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	if balance < -1 && getBalance(node.right) <= 0 {
		return rotateLeft(node)
	}
	if balance < -1 && getBalance(node.right) > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	return node
}

func height[K cmp.Ordered, V any](node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance[K cmp.Ordered, V any](node *Node[K, V]) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func getMinNode[K cmp.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func rotateRight[K cmp.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	newRoot := node.left
	t2 := newRoot.right

	newRoot.right = node
	node.left = t2

	node.height = max(height(node.left), height(node.right)) + 1
	newRoot.height = max(height(newRoot.left), height(newRoot.right)) + 1

	return newRoot
}

func rotateLeft[K cmp.Ordered, V any](node *Node[K, V]) *Node[K, V] {
	newRoot := node.right
	t2 := newRoot.left

	newRoot.left = node
	node.right = t2

	node.height = max(height(node.left), height(node.right)) + 1
	newRoot.height = max(height(newRoot.left), height(newRoot.right)) + 1

	return newRoot
}
