package linear

type ArrayStack struct {
	list *ArrayList
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		list: NewArrayList(),
	}
}

func (s *ArrayStack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *ArrayStack) Push(value int) bool {
	return s.list.PushBack(value)
}

func (s *ArrayStack) Pop() (int, bool) {
	return s.list.DeleteAt(s.list.Size() - 1)
}
