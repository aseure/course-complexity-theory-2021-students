package solvers

import "github/aseure/course-complexity-theory-2021/maze"

func BacktrackingSolver(start, end *maze.Cell) {
	start.SetMark(maze.Visiting)
	defer start.SetMark(maze.Path)
	_ = backtrackingSolver(start, end)
}

func backtrackingSolver(start, end *maze.Cell) bool {
	if start == end {
		return true
	}

	for d, neighbor := range start.GetRandomizedNeighbors() {
		if neighbor != nil && neighbor.GetMark() != maze.Visiting && !start.IsWall(d) {
			neighbor.SetMark(maze.Visiting)

			if ok := backtrackingSolver(neighbor, end); ok {
				neighbor.SetMark(maze.Path)
				return true
			}

			neighbor.SetMark(maze.Blank)
		}
	}

	return false
}
