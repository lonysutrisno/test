package algo

type Stack []string

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(arg string) {
	*s = append(*s, arg)
}

func (s *Stack) Pop() (element string) {
	if s.IsEmpty() {
		return
	}
	index := len(*s) - 1
	element = (*s)[index]
	*s = (*s)[:index]
	return element
}

// func main() {

// 	var stack Stack // create a stack variable of type Stack

// 	stack.Push("this")
// 	stack.Push("is")
// 	stack.Push("sparta!!")

// 	for len(stack) > 0 {
// 		x := stack.Pop()
// 		fmt.Println(x)
// 	}
// }
