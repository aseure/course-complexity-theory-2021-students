package maze

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMaze(t *testing.T) {
	m := NewMaze(100, 100, 2, 20)
	m.Generate()

	start := time.Now()
	m.BacktrackingSolver()
	fmt.Printf("Backtracking solver took %s\n", time.Since(start))
	require.NoError(t, m.Display())

	start = time.Now()
	m.Reset()
	m.DijkstraSolver()
	fmt.Printf("Dijkstra solver took %s\n", time.Since(start))
	require.NoError(t, m.Display())
}
