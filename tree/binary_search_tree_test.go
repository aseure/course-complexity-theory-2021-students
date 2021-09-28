package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	require.Equal(t, 0, bst.Size())

	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(4)
	bst.Insert(3)
	require.Equal(t, 4, bst.Size())

	require.True(t, bst.Search(1))
	require.True(t, bst.Search(2))
	require.True(t, bst.Search(3))
	require.True(t, bst.Search(4))

	require.True(t, bst.Remove(1))
	require.Equal(t, 3, bst.Size())
	require.True(t, bst.Remove(2))
	require.Equal(t, 2, bst.Size())
	require.True(t, bst.Remove(3))
	require.Equal(t, 1, bst.Size())
	require.True(t, bst.Remove(4))
	require.Equal(t, 0, bst.Size())
	require.False(t, bst.Remove(5))
	require.Equal(t, 0, bst.Size())

}
