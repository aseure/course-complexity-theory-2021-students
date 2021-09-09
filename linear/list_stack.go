package linear

type ListStack struct {
	list *DoublyLinkedList
}

func NewListStack() *ListStack {
	return &ListStack{
		list: NewDoublyLinkedList(),
	}
}

func (s *ListStack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *ListStack) Push(value int) bool {
	return s.list.PushFront(value)
}

func (s *ListStack) Pop() (int, bool) {
	return s.list.DeleteAt(0)
}
