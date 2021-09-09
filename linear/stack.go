package linear

type Stack interface {
	IsEmpty() bool
	Push(value int) bool
	Pop() (int, bool)
}
