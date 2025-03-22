package parser

type binaryTreeValue interface {
	lessOrEqualInPrecedence(op Symbol) bool
}
type BinaryTreeNode struct {
	Value binaryTreeValue
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}
