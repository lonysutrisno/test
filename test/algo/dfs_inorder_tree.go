package algo

import "fmt"

func DFSInorder(n *Node) {
	if n == nil {
		return
	}

	if !existInResult(n.data) {

		DFSInorder(n.left)
		result = append(result, n.data)
		DFSInorder(n.right)
	}
	l, r := n.Next()
	if l == nil && r == nil {
		return
	}

}

func ExecDFSTreeInorder() {
	tree := ExecTree()
	DFSInorder(tree.root)
	fmt.Println(result)
}
