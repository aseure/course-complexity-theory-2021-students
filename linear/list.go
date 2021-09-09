package linear

type List interface {
	Size() int
	Values() []int
	GetAt(position int) (int, bool)
	InsertAt(position, value int) bool
	DeleteAt(position int) (int, bool)
	PushBack(value int) bool
	PushFront(value int) bool
}
