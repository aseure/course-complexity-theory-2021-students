package tree

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (t *BinarySearchTree) Size() int {
	return t.root.Size()
}

func (t *BinarySearchTree) Height() int {
	return t.root.Height()
}

func (t *BinarySearchTree) Insert(value int) {
	t.root = t.root.Insert(value)
}

func (t *BinarySearchTree) Remove(value int) bool {
	var ok bool
	t.root, ok = t.root.Remove(value)
	return ok
}

func (t *BinarySearchTree) Search(value int) bool {
	return t.root.Search(value)
}

func (t *BinarySearchTree) Display() error {
	start := func(w func(format string, a ...interface{})) {
		visitor := func(n *BinarySearchTreeNode) {
			if n.left != nil {
				w("  %d -> %d;\n", n.value, n.left.value)
			}
			if n.right != nil {
				w("  %d -> %d;\n", n.value, n.right.value)
			}
		}

		if t.root != nil {
			if t.root.left == nil && t.root.right == nil {
				w("  %d;\n", t.root.value)
			} else {
				t.root.Accept(visitor)
			}
		}
	}

	return display(start)
}
