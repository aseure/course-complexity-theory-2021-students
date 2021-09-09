package linear

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

type SinglyLinkedListNode struct {
	value int
	next  *SinglyLinkedListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
	}
}

func (l *SinglyLinkedList) Size() int {
	size := 0
	node := l.head
	for node != nil {
		size++
		node = node.next
	}
	return size
}

func (l *SinglyLinkedList) Values() []int {
	var values []int
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *SinglyLinkedList) GetAt(position int) (int, bool) {
	node := l.head
	for i := 0; i < position && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return 0, false
	}
	return node.value, true
}

func (l *SinglyLinkedList) InsertAt(position, value int) bool {
	if position == 0 {
		return l.PushFront(value)
	}
	if l.head == nil {
		return false
	}
	prev := l.head
	for i := 1; i < position && prev != nil; i++ {
		prev = prev.next
	}
	if prev == nil {
		return false
	}
	newNode := &SinglyLinkedListNode{
		value: value,
		next:  prev.next,
	}
	prev.next = newNode
	return true
}

func (l *SinglyLinkedList) DeleteAt(position int) (int, bool) {
	if l.head == nil || position < 0 {
		return 0, false
	}
	if position == 0 {
		value := l.head.value
		l.head = l.head.next
		return value, true
	}
	prev := l.head
	for i := 1; i < position && prev != nil; i++ {
		prev = prev.next
	}
	if prev == nil {
		return 0, false
	}
	if prev.next == nil {
		return 0, false
	}
	value := prev.next.value
	prev.next = prev.next.next
	return value, true
}

func (l *SinglyLinkedList) PushBack(value int) bool {
	newNode := &SinglyLinkedListNode{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		return true
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
	return true
}

func (l *SinglyLinkedList) PushFront(value int) bool {
	newNode := &SinglyLinkedListNode{
		value: value,
		next:  l.head,
	}
	l.head = newNode
	return true
}
