package ds

type Node[T any] struct {
    value T
    next *Node[T]
}

type LinkedList[T any] struct {
    head *Node[T]
}
