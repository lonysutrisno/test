package hrank

import "fmt"

func ExecReversedLinkedlist() {
	param := []int{1, 2, 3}
	var linked List
	for _, j := range param {
		linked.InsertNormal(j)
	}
	ReversedLinkedlist(linked.Head)
	n := linked.Head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}

}
func ReversedLinkedlist(n *Node) {
	var l *Node
	l = n
	for n != nil {
		l = insert(l, n.value)
		n = n.next
	}

}
func insert(l *Node, data int) *Node {

	list := &Node{value: data}
	if l == nil {
		l = list
	} else {
		p := l
		l = nil
		l = &Node{value: data, next: p}
	}
	return l
}
