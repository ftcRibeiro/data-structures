package datastructure

import "golang.org/x/exp/constraints"

type Slice[T any] []T

type Ordered interface {
	constraints.Ordered
}

type Number interface {
	constraints.Integer | constraints.Float
}
