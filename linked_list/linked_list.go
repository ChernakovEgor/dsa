package linkedlist

type node[T any] struct {
	v    T
	next *node[T]
}

type LinkedList[T any] struct {
	root *node[T]
	tail node[T]
}

func (l LinkedList[T]) Len() int
func (l *LinkedList[T]) InsertAt(int)
func (l *LinkedList[T]) Remove(T) T
func (l *LinkedList[T]) RemoveAt(int) T
func (l *LinkedList[T]) Append(T)
func (l *LinkedList[T]) Prepend(T)
func (l LinkedList[T]) Get(int) T
