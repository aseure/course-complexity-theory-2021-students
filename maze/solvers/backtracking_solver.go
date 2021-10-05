package solvers

import "github/aseure/course-complexity-theory-2021/maze"

func BacktrackingSolver(m *maze.Maze, start, end *maze.Cell) {
	start.SetMark(maze.Visiting)
	_ = backtrackingSolver(start, end)
	start.SetMark(maze.Path)
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
