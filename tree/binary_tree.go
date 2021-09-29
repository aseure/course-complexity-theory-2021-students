package tree

type BinaryTree interface {
	Size() int
	Height() int
	Insert(value int)
	Remove(value int) bool
	Search(value int) bool
	Display() error
}
