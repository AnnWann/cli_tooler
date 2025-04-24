package utility

type BinaryTreeNode[T any] struct {
	Value T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}
