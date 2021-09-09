package linear

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

type DoublyLinkedListNode struct {
	value int
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

func (l *DoublyLinkedList) Size() int {
	size := 0
	node := l.head
	for node != nil {
		size++
		node = node.next
	}
	return size
}

func (l *DoublyLinkedList) Values() []int {
	var values []int
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *DoublyLinkedList) GetAt(position int) (int, bool) {
	node := l.head
	for i := 0; i < position && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return 0, false
	}
	return node.value, true
}

func (l *DoublyLinkedList) InsertAt(position, value int) bool {
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
	newNode := &DoublyLinkedListNode{
		value: value,
		prev:  prev,
		next:  prev.next,
	}
	if prev.next != nil {
		prev.next.prev = newNode
	}
	prev.next = newNode
	return true
}

func (l *DoublyLinkedList) DeleteAt(position int) (int, bool) {
	if l.head == nil || position < 0 {
		return 0, false
	}
	if position == 0 {
		value := l.head.value
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
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
	if prev.next.next != nil {
		prev.next.next.prev = prev
	}
	prev.next = prev.next.next
	return value, true
}

func (l *DoublyLinkedList) PushBack(value int) bool {
	newNode := &DoublyLinkedListNode{
		value: value,
		prev:  l.tail,
		next:  nil,
	}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return true
	}
	l.tail.next = newNode
	l.tail = newNode
	return true
}

func (l *DoublyLinkedList) PushFront(value int) bool {
	newNode := &DoublyLinkedListNode{
		value: value,
		prev:  nil,
		next:  l.head,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return true
	}
	l.head.prev = newNode
	l.head = newNode
	return true
}
