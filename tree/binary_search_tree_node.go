package tree

type BinarySearchTreeNode struct {
	value int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(value int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (n *BinarySearchTreeNode) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *BinarySearchTreeNode) Height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.Height(), n.right.Height())
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (n *BinarySearchTreeNode) Insert(value int) *BinarySearchTreeNode {
	if n == nil {
		return NewBinarySearchTreeNode(value)
	}
	if value < n.value {
		n.left = n.left.Insert(value)
	}
	if value > n.value {
		n.right = n.right.Insert(value)
	}
	return n
}

func (n *BinarySearchTreeNode) Remove(value int) (*BinarySearchTreeNode, bool) {
	if n == nil {
		return nil, false
	}
	var ok bool
	if value < n.value {
		n.left, ok = n.left.Remove(value)
		return n, ok
	}
	if value > n.value {
		n.right, ok = n.right.Remove(value)
		return n, ok
	}
	if n.left != nil {
		n.left, n.value = n.left.removeMax()
		return n, true
	}
	if n.right != nil {
		n.right, n.value = n.right.removeMin()
		return n, true
	}
	return nil, true
}

func (n *BinarySearchTreeNode) Search(value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return n.left.Search(value)
	}
	if value > n.value {
		return n.right.Search(value)
	}
	return true
}

func (n *BinarySearchTreeNode) removeMin() (*BinarySearchTreeNode, int) {
	if n.left == nil {
		return nil, n.value
	}
	var min int
	n.left, min = n.left.removeMin()
	return n, min
}

func (n *BinarySearchTreeNode) removeMax() (*BinarySearchTreeNode, int) {
	if n.right == nil {
		return nil, n.value
	}
	var max int
	n.right, max = n.right.removeMax()
	return n, max
}

func (n *BinarySearchTreeNode) Accept(v BinarySearchTreeNodeVisitor) {
	if n == nil {
		return
	}
	n.left.Accept(v)
	v(n)
	n.right.Accept(v)
}
