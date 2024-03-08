package algo

import "fmt"

func DFSPostorder(n *Node) {
	if n == nil {
		return
	}

	if !existInResult(n.data) {

		DFSPostorder(n.left)
		DFSPostorder(n.right)
		result = append(result, n.data)

	}
	l, r := n.Next()
	if l == nil && r == nil {
		return
	}

}

func ExecDFSTreePostorder() {
	tree := ExecTree()
	DFSPostorder(tree.root)
	fmt.Println(result)
}
