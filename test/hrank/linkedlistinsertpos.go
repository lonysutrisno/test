package hrank

import "fmt"

func insertNodeAtPosition(llist *Node, data int, position int32) *Node {
	// Write your code here
	p := llist
	for i := 0; i < int(position); i++ {
		p = p.next
	}
	temp := p.next
	p.next = &Node{value: data, next: temp}
	return llist
}

func ExecInsertAtPosition() {
	param := []int{1, 2, 3}
	var linked List
	for _, j := range param {
		linked.InsertHead(j)
	}
	insertNodeAtPosition(linked.Head, 5, 1)
	n := linked.Head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}
