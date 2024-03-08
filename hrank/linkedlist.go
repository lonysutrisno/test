package hrank

import (
	"fmt"
)

func ExecLinkedlist() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// reader := bufio.NewReader(os.Stdin)
	// var n, input string
	// fmt.Scanf("%s", &n)
	// var param []int
	// i, _ := strconv.Atoi(n)
	// for i > 0 {
	// 	fmt.Scanf("%s", &input)
	// 	inputInt, _ := strconv.Atoi(input)
	// 	param = append(param, inputInt)
	// 	i--
	// }
	param := []int{1, 2, 3}
	var linked List
	for _, j := range param {
		linked.InsertNormal(j)

	}
	n := linked.Head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}

}

func (l *List) InsertNormal(data int) {
	list := &Node{value: data}
	if l.Head == nil {
		l.Head = list
	} else {
		p := l.Head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}
