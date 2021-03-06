package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryTrees(t *testing.T) {
	for _, c := range []struct {
		description string
		tree        BinaryTree
	}{
		{"binary search tree", NewBinarySearchTree()},
		{"AVL tree", NewAVLTree()},
	} {
		t.Run(c.description, func(t *testing.T) {
			require.Equal(t, 0, c.tree.Size())

			nbInserts := 10
			for i := 1; i <= nbInserts; i++ {
				c.tree.Insert(i)
			}
			require.Equal(t, nbInserts, c.tree.Size())

			if tree, ok := c.tree.(*AVLTree); ok {
				fmt.Printf("AVLTree DFS: %v\n", tree.DepthFirstTraversal())
				fmt.Printf("AVLTree BFS: %v\n", tree.BreadthFirstTraversal())
			}

			for i := 1; i <= nbInserts; i++ {
				require.True(t, c.tree.Search(i))
				require.True(t, c.tree.Remove(i))
				require.False(t, c.tree.Search(i))
				require.Equal(t, nbInserts-i, c.tree.Size())
			}
		})
	}
}
