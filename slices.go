package datastructure

import "sort"

func RemoveByItem[T comparable](s Slice[T], item T) Slice[T] {
	for i, v := range s {
		if v == item {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func RemoveByIndex[T comparable](s Slice[T], i int) Slice[T] {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func SumItems[T Number](s Slice[T]) T {
	var sum T
	for _, v := range s {
		sum = sum + v
	}
	return sum
}

func OrderItems[T Ordered](s []T) []T {
	var ordered []T
	for _, v := range s {
		ordered = append(ordered, v)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i] < ordered[j]
	})
	return ordered
}
