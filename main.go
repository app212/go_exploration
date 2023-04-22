package main

import "fmt"

func main() {
	list := NewLinkedList(10)
	fmt.Println(list.FirstNode.Value, list.Length)
	list.AddLast(21)
	fmt.Println(list.FirstNode.Next.Value, list.Length)

	for i := 0; i < 10; i++ {
		list.AddLast(i * 5)
	}

	fmt.Println(list)

	list.AddFirst(100)
	fmt.Println(list)

	list.RemoveFirst()
	fmt.Println(list)

	list.RemoveLast()
	fmt.Println(list)

	list2 := NewLinkedList("abc")
	fmt.Println(list2)

	list2.RemoveFirst()
	fmt.Println(list2)

	list2.AddFirst("World")
	fmt.Println(list2)

	list2.AddFirst("Hello")
	fmt.Println(list2)

	node1 := list2.GetNode(3)
	fmt.Println(node1)

	node2 := list2.GetNode(0)
	fmt.Println(node2)

	node3 := list2.GetNode(1)
	fmt.Println(node3)

	fmt.Println(list.GetNode(10))

	list2.AddAtPos("Abc", 0)
	list2.AddAtPos("Abc2", 5)
	list2.AddAtPos("Abc3", 2)
	list2.AddAtPos("Abc4", 4)
	fmt.Println(list2)

	list2.RemoveAtPos(0)
	list2.RemoveAtPos(1)
	list2.RemoveAtPos(2)
	list2.RemoveAtPos(6)
	fmt.Println(list2)
}
