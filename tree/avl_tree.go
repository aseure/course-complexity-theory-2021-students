package tree

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func (t *AVLTree) Size() int {
	return t.root.Size()
}

func (t *AVLTree) Height() int {
	return t.root.Height()
}

func (t *AVLTree) Insert(value int) {
	t.root = t.root.Insert(value)
}

func (t *AVLTree) Remove(value int) bool {
	var ok bool
	t.root, ok = t.root.Remove(value)
	return ok
}

func (t *AVLTree) Search(value int) bool {
	return t.root.Search(value)
}

func (t *AVLTree) Display() error {
	start := func(w func(format string, a ...interface{})) {
		visitor := func(n *AVLTreeNode) {
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

func (t *AVLTree) DepthFirstTraversal() (values []int) {
	if t.root == nil {
		return
	}
	t.root.Accept(func(n *AVLTreeNode) {
		values = append(values, n.value)
	})
	return
}

func (t *AVLTree) BreadthFirstTraversal() (values []int) {
	if t.root == nil {
		return
	}

	queue := []*AVLTreeNode{t.root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n != nil {
			queue = append(queue, n.left, n.right)
			values = append(values, n.value)
		}
	}

	return
}
