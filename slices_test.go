package datastructure

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveByItem(t *testing.T) {
	t.Run("RemoveByItem", func(t *testing.T) {
		test := []struct {
			s        Slice[int]
			item     int
			expected Slice[int]
		}{
			{Slice[int]{1, 2, 3, 4, 5}, 3, Slice[int]{1, 2, 4, 5}},
			{Slice[int]{1, 2, 3, 4, 5}, 4, Slice[int]{1, 2, 3, 5}},
			{Slice[int]{1, 2, 3, 4, 5}, 5, Slice[int]{1, 2, 3, 4}},
			{Slice[int]{1, 2, 3, 4, 5}, 6, Slice[int]{1, 2, 3, 4, 5}},
		}
		for _, v := range test {
			actual := RemoveByItem(v.s, v.item)
			require.Equal(t, v.expected, actual)
		}
	})
}

func TestRemoveByIndex(t *testing.T) {
	t.Run("RemoveByIndex", func(t *testing.T) {
		test := []struct {
			s        Slice[int]
			i        int
			expected Slice[int]
			err      error
		}{
			{Slice[int]{1, 2, 3, 4, 5}, 0, Slice[int]{2, 3, 4, 5}, nil},
			{Slice[int]{1, 2, 3, 4, 5}, 2, Slice[int]{1, 2, 4, 5}, nil},
			{Slice[int]{1, 2, 3, 4, 5}, 1, Slice[int]{1, 3, 4, 5}, nil},
			{Slice[int]{1, 2, 3, 4, 5}, 3, Slice[int]{1, 2, 3, 5}, nil},
			{Slice[int]{1, 2, 3, 4, 5}, 5, nil, errors.New("index out of range")},
		}
		for _, v := range test {
			actual, err := RemoveByIndex(v.s, v.i)
			require.Equal(t, v.err, err)
			require.Equal(t, v.expected, actual)
		}
	})
}

func TestSumItems(t *testing.T) {
	t.Run("SumItems", func(t *testing.T) {
		test := []struct {
			s        Slice[int]
			expected int
		}{
			{Slice[int]{1, 2, 3, 4, 5}, 15},
			{Slice[int]{5, 6, 3, 4, 5}, 23},
			{Slice[int]{1, 5, 3, 4, 8}, 21},
		}
		for _, v := range test {
			actual := SumItems(v.s)
			require.Equal(t, v.expected, actual)
		}
	})

}

func TestOrderItems(t *testing.T) {
	t.Run("OrderItems", func(t *testing.T) {
		test := []struct {
			s        []int
			expected []int
		}{
			{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			{[]int{5, 6, 3, 4, 5}, []int{3, 4, 5, 5, 6}},
			{[]int{1, 5, 3, 4, 8}, []int{1, 3, 4, 5, 8}},
		}
		for _, v := range test {
			actual := OrderItems(v.s)
			require.Equal(t, v.expected, actual)
		}
	})
}
