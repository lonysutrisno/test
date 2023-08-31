package algo

import "fmt"

func DFS(n *Node) {
	if n == nil {
		return
	}

	if !existInResult(n.data) {
		result = append(result, n.data)
		DFS(n.left)
		DFS(n.right)
	}
	l, r := n.Next()
	if l == nil && r == nil {
		return
	}

}

func existInResult(x int64) bool {
	for _, j := range result {
		if j == x {
			return true
		}
	}
	return false
}
func (n *Node) Next() (left *Node, right *Node) {
	return n.left, n.right
}

func ExecDFSTree() {
	tree := ExecTree()
	DFS(tree.root)
	fmt.Println(result)
}
