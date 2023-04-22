package main

import "fmt"

// TODO: error handling

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	FirstNode *Node[T]
	Length    int
}

func NewLinkedList[T any](val T) LinkedList[T] {
	node := Node[T]{val, nil}
	return LinkedList[T]{&node, 1}
}

func (list *LinkedList[T]) AddLast(val T) {
	if list.FirstNode == nil {
		list.FirstNode = &Node[T]{val, nil}
		list.Length += 1
	} else {
		currentNode := list.FirstNode
		nextNode := currentNode.Next

		for nextNode != nil {
			currentNode = nextNode
			nextNode = currentNode.Next
		}

		currentNode.Next = &Node[T]{val, nil}
		list.Length += 1
	}
}

func (list *LinkedList[T]) AddFirst(val T) {
	if list.FirstNode == nil {
		list.FirstNode = &Node[T]{val, nil}
		list.Length += 1
	} else {
		oldFirst := list.FirstNode
		list.FirstNode = &Node[T]{val, oldFirst}
		list.Length += 1
	}
}

func (list *LinkedList[T]) RemoveFirst() {
	if list.FirstNode == nil {
		return
	}

	list.FirstNode = list.FirstNode.Next
	list.Length -= 1
}

func (list *LinkedList[T]) RemoveLast() {
	if list.FirstNode == nil {
		return
	}

	currentNode := list.FirstNode
	nextNode := currentNode.Next
	for nextNode.Next != nil {
		currentNode = nextNode
		nextNode = currentNode.Next
	}
	currentNode.Next = nil
	list.Length -= 1
}

func (list *LinkedList[T]) GetNode(n int) *Node[T] {
	if list.Length <= n {
		return nil
	}
	index := 0
	current := list.FirstNode
	for index < n {
		current = current.Next
		index++
	}
	return current
}

func (list *LinkedList[T]) AddAtPos(val T, pos int) {
	if pos == 0 {
		list.AddFirst(val)
	} else if pos > list.Length {
		// error
		return
	} else {
		oldNode := list.GetNode(pos)
		prevNode := list.GetNode(pos - 1)

		newNode := &Node[T]{val, oldNode}
		prevNode.Next = newNode
		list.Length++
	}
}

func (list *LinkedList[T]) RemoveAtPos(pos int) {
	if pos == 0 {
		list.RemoveFirst()
	} else if pos >= list.Length {
		// error
		return
	} else {
		oldNode := list.GetNode(pos)
		prevNode := list.GetNode(pos - 1)

		prevNode.Next = oldNode.Next
		oldNode.Next = nil
		list.Length--
	}
}

func (list LinkedList[T]) String() string {
	if list.FirstNode == nil {
		return "Empty LinkedList"
	}
	var values string = "[" + fmt.Sprint(list.FirstNode.Value) + "]"
	currentNode := list.FirstNode
	nextNode := currentNode.Next

	for nextNode != nil {
		currentNode = nextNode
		nextNode = currentNode.Next
		value := "[" + fmt.Sprint(currentNode.Value) + "]"
		values += value
	}

	return fmt.Sprintf("Length: %d, Values: %s", list.Length, values)
}

func (node Node[T]) String() string {
	return fmt.Sprintf("Node: %v", node.Value)
}
