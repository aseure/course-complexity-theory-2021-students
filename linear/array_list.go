package linear

type ArrayList struct {
	arr      []int
	capacity int
	end      int
}

func NewArrayList() *ArrayList {
	capacity := 10
	return &ArrayList{
		arr:      make([]int, capacity),
		capacity: capacity,
		end:      0,
	}
}

func (l *ArrayList) Size() int {
	return l.end
}

func (l *ArrayList) Values() []int {
	return l.arr[:l.end]
}

func (l *ArrayList) GetAt(position int) (int, bool) {
	if position >= l.end {
		return 0, false
	}
	return l.arr[position], true
}

func (l *ArrayList) InsertAt(position, value int) bool {
	if position < 0 || l.end < position {
		return false
	}
	if l.end == l.capacity {
		l.growCapacity()
	}
	for i := l.end; i > position; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[position] = value
	l.end++
	return true
}

func (l *ArrayList) DeleteAt(position int) (int, bool) {
	if position < 0 || l.end < position {
		return 0, false
	}
	value := l.arr[position]
	for i := position; i < l.end-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.end--
	return value, true
}

func (l *ArrayList) PushBack(value int) bool {
	if l.end == l.capacity {
		l.growCapacity()
	}
	l.arr[l.end] = value
	l.end++
	return true
}

func (l *ArrayList) PushFront(value int) bool {
	if l.end == l.capacity {
		l.growCapacity()
	}
	for i := l.end; i > 0; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[0] = value
	l.end++
	return true
}

func (l *ArrayList) growCapacity() {
	l.capacity = l.capacity * 2
	newArr := make([]int, l.capacity)
	copy(newArr, l.arr)
	l.arr = newArr
}
