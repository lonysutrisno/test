package script

import "fmt"

type LinkedList struct {
	Head *Node
}
type Node struct {
	Value int64
	Next  *Node
}

func NewLinkList() *LinkedList {
	return &LinkedList{}
}
func (ll *LinkedList) Insert(next Node) {
	if ll.Head == nil {
		ll.Head = &next
	} else {

		temp := ll.Head
		next.Next = temp
		ll.Head = &next

		// for temp.Next != nil {
		// 	temp = temp.Next
		// }
		// temp.Next = &next
	}

}

func (ll *LinkedList) Reverse() LinkedList {
	var newLL LinkedList
	if ll.Head == nil {
		fmt.Println("kosong")
	} else {
		temp := ll.Head
		for temp.Next != nil {
			newLL.Insert(*temp)
			temp = temp.Next
		}

	}
	return newLL
}

func (ll *LinkedList) Print() {
	if ll.Head == nil {
		fmt.Println("empty")
	} else {
		fmt.Println(ll.Head.Value)
		temp := ll.Head
		for temp.Next != nil {
			temp = temp.Next
			fmt.Println(temp.Value)
		}

	}

}
func InitLL() {
	ll := NewLinkList()
	ll.Insert(Node{Value: 12})
	ll.Insert(Node{Value: 13})
	ll.Insert(Node{Value: 15})
	ll.Insert(Node{Value: 18})
	ll.Print()
	fmt.Println("asd")
	rev := ll.Reverse()
	rev.Print()
}
