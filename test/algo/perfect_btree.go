package algo

import (
	"strconv"
)

type Node struct {
	left   *Node
	right  *Node
	data   int64
	parent *Node
}

type Tree struct {
	root *Node
}

var LevelNode = make(map[int][]*Node)
var Level int

var result []int64

func AppendLevelNode(lv int, node *Node) {
	LevelNode[lv] = append(LevelNode[lv], node)
}

// func ReplaceNode(n *Node) {

// 	for idx, node := range LevelNode[Level] {
// 		if n == node {
// 			fmt.Println(n)
// 			fmt.Println(node)
// 			LevelNode[Level][idx] = n
// 		}
// 	}
// }

func (t *Tree) Insert(data int64) {
	if t.IsEmpty() {
		t.root = &Node{left: nil, right: nil, data: data}
		AppendLevelNode(Level, t.root)

	} else {
		//if empty children
		if t.root.left == nil || t.root.right == nil {
			Level = 1
			if t.root.left == nil {
				t.root.left = &Node{left: nil, right: nil, data: data, parent: t.root}
				AppendLevelNode(Level, t.root.left)

			} else if t.root.right == nil {
				t.root.right = &Node{left: nil, right: nil, data: data, parent: t.root}
				AppendLevelNode(Level, t.root.right)
			}
		} else {
			t.root.left.InsertNode(data)
		}
	}
}

func (t *Tree) IsEmpty() bool {
	return t.root == nil
}

func (n *Node) CheckEmptyNode(data int64) (EmptyNode *Node) {

	if Level > 1 {
		lv := Level

		for _, node := range LevelNode[lv-1] {
			if node.left != nil && node.right != nil {
				continue
			} else if node.left == nil || node.right == nil {
				return node
			}
		}
	}
	for _, node := range LevelNode[Level] {
		if node.left != nil && node.right != nil {
			continue
		} else if node.left == nil || node.right == nil {
			return node
		}
	}

	return nil
}

func (n *Node) InsertNode(data int64) {
	EmptyNode := n.CheckEmptyNode(data)
	if EmptyNode == nil {
		//new level
		n.left = &Node{left: nil, right: nil, data: data, parent: n}
		Level++
		AppendLevelNode(Level, n.left)
	} else if EmptyNode.left == nil {
		EmptyNode.left = &Node{left: nil, right: nil, data: data, parent: EmptyNode}
		AppendLevelNode(Level, EmptyNode.left)
	} else if EmptyNode.right == nil {
		EmptyNode.right = &Node{left: nil, right: nil, data: data, parent: EmptyNode}
		AppendLevelNode(Level, EmptyNode.right)
	}

}
func binaryTreePathsUtil(root *Node, curr string, output *[]string) {
	if root == nil {
		return
	}

	valString := strconv.Itoa(int(root.data))
	if curr == "" {
		curr = valString
	} else {
		curr = curr + "->" + valString
	}
	if root.left == nil && root.right == nil {
		*output = append(*output, curr)
		return
	}

	binaryTreePathsUtil(root.left, curr, output)
	binaryTreePathsUtil(root.right, curr, output)

}
func binaryTreePaths(root *Node) []string {
	output := make([]string, 0)

	binaryTreePathsUtil(root, "", &output)
	return output
}
func ExecTree() Tree {
	var t Tree

	for i := 1; i <= 10; i++ {
		t.Insert(int64(i))
	}
	// fmt.Println(binaryTreePaths(t.root))
	// fmt.Println(LevelNode)
	return t
}
