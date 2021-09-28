package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	require.Equal(t, 0, h.Size())

	h.Insert(2)
	h.Insert(19)
	h.Insert(100)
	h.Insert(3)
	h.Insert(36)
	h.Insert(7)
	h.Insert(17)
	h.Insert(25)
	h.Insert(1)
	require.Equal(t, 9, h.Size())

	max, ok := h.PopMax()
	require.True(t, ok)
	require.Equal(t, 100, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 36, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 25, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 19, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 17, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 7, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 3, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 2, max)

	max, ok = h.PopMax()
	require.True(t, ok)
	require.Equal(t, 1, max)

	require.Equal(t, 0, h.Size())

	_, ok = h.PopMax()
	require.False(t, ok)
}
