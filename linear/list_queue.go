package linear

type ListQueue struct {
	list *DoublyLinkedList
}

func NewListQueue() *ListQueue {
	return &ListQueue{
		list: NewDoublyLinkedList(),
	}
}

func (q *ListQueue) IsEmpty() bool {
	return q.list.Size() == 0
}

func (q *ListQueue) Push(value int) bool {
	return q.list.PushBack(value)
}

func (q *ListQueue) Pop() (int, bool) {
	return q.list.DeleteAt(0)
}
