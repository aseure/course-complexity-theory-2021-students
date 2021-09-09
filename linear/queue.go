package linear

type Queue interface {
	IsEmpty() bool
	Push(value int) bool
	Pop() (int, bool)
}
