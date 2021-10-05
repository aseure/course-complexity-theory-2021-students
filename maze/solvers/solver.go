package solvers

import "github/aseure/course-complexity-theory-2021/maze"

type Solver func(m *maze.Maze, start, end *maze.Cell)
