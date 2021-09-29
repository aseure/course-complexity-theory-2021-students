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

type BinarySearchTreeNodeVisitor func(n *BinarySearchTreeNode)

func (n *BinarySearchTreeNode) Accept(v BinarySearchTreeNodeVisitor) {
	if n == nil {
		return
	}
	n.left.Accept(v)
	v(n)
	n.right.Accept(v)
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
	} else if value > n.value {
		n.right, ok = n.right.Remove(value)
		return n, ok
	} else {
		if n.left != nil {
			n.value = n.left.getMax()
			n.left, ok = n.left.Remove(n.value)
			return n, ok
		} else if n.right != nil {
			n.value = n.right.getMin()
			n.right, ok = n.right.Remove(n.value)
			return n, ok
		} else {
			return nil, true
		}
	}
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

func (n *BinarySearchTreeNode) getMin() int {
	if n.left == nil {
		return n.value
	}
	return n.left.getMin()
}

func (n *BinarySearchTreeNode) getMax() int {
	if n.right == nil {
		return n.value
	}
	return n.right.getMax()
}
