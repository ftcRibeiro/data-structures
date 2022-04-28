package datastructure

import (
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
		}
		for _, v := range test {
			actual := RemoveByItem(v.s, v.item)
			require.Equal(t, v.expected, actual)
		}
	})
}
