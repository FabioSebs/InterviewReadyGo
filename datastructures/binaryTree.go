package datastructures

// binary tree is a heirarchical data structure
// each node has a left and right child

type BNode struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *BNode {
	return &BNode{Value: value, Left: nil, Right: nil}
}

/*
types of searches
	- Preorder
	- Inorder
	- Postorder
	- Breadth First Search
*/
