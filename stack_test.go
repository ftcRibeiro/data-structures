package datastructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewStack(t *testing.T) {
	stack := NewStack[string]()
	require.Nil(t, stack.Root)
}

func Test_Functions(t *testing.T) {
	stack := NewStack[int]()

	t.Run("Empty", func(t *testing.T) {
		require.True(t, stack.Empty())
	})

	t.Run("Pop Empty", func(t *testing.T) {
		item := stack.Pop()
		require.Nil(t, item)
	})

	t.Run("Push", func(t *testing.T) {
		stack.Push(1)
		require.Equal(t, 1, stack.Root.Val)
	})

	t.Run("Pop", func(t *testing.T) {
		item := stack.Pop()
		require.Equal(t, 1, *item)
	})

}
