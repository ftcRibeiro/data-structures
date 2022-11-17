package datastructure

import "golang.org/x/exp/constraints"

type Slice[T any] []T

type Ordered interface {
	constraints.Ordered
}

type Number interface {
	constraints.Integer | constraints.Float
}

// type TreeNode[T Item] interface {
// 	Esq() *TreeNode[T]
// 	Dir() *TreeNode[T]
// 	Value() *T
// }

// type Tree[T any] interface {
// 	Root() *TreeNode[T]
// }

type Node[T any] struct {
	Next *Node[T]
	Val  T
}

type Stack[T any] struct {
	Root *Node[T]
}

// type Item interface {
// 	// Retrives string value
// 	Value() string
// 	// Retrives string value
// 	SetValue(v string)
// 	// Retrives indentifier
// 	ID() string
// 	// Retrives comparable hash
// 	Hash() string
// 	// Retrives left child
// 	Left() *Item
// 	// Retrives right child
// 	Right() *Item
// }
