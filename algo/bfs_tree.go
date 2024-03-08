package algo

import "fmt"

func BFS(n *Node) {
	if n == nil {
		return
	}

	if len(result) == 0 {
		result = append(result, n.data)
	} else {
		l, r := n.Next()
		if l == nil && r == nil {
			return
		} else {
			result = append(result, n.left.data, n.right.data)
		}
	}
	BFS(n.left)
	BFS(n.right)

}

func ExecBFSTree() {
	tree := ExecTree()
	BFS(tree.root)
	fmt.Println(result)
}
