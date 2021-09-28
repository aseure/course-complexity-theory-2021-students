package tree

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

func (t *BinarySearchTree) GenerateDot(filename string) error {
	var b strings.Builder

	w := func(format string, a ...interface{}) {
		b.WriteString(fmt.Sprintf(format, a...))
	}

	w("digraph bst{\n")
	if t.root != nil {
		if t.root.left == nil && t.root.right == nil {
			w("  %d;\n", t.root.value)
		} else {
			t.root.Accept(func(n *BinarySearchTreeNode) {
				if n.left != nil {
					w("  %d -> %d;\n", n.value, n.left.value)
				}
				if n.right != nil {
					w("  %d -> %d;\n", n.value, n.right.value)
				}
			})
		}
	}
	w("}\n")

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("cannot open Dot file %q for writing: %w", filename, err)
	}

	cmd := exec.Command("dot", "-Tsvg")
	cmd.Stdin = strings.NewReader(b.String())
	cmd.Stdout = f

	errCmd := cmd.Run()
	errClose := f.Close()

	if errCmd != nil {
		return fmt.Errorf("could not generate Dot file %q correctly: %w", filename, errCmd)
	}

	if errClose != nil {
		return fmt.Errorf("could not close Dot file %q correctly: %w", filename, errClose)
	}

	return nil
}
