package maze

func (m *Maze) BacktrackingSolver() {
	start, end := m.cells[0][0], m.cells[m.height-1][m.width-1]
	visited := make(map[*Cell]bool)
	start.isPath = backtrackingSolver(visited, start, end)
}

func backtrackingSolver(visited map[*Cell]bool, start, end *Cell) bool {
	if start == end {
		return true
	}

	visited[start] = true

	for d, neighbor := range start.GetRandomizedNeighbors() {
		if neighbor != nil && !visited[neighbor] && !start.isWall(d) {
			if ok := backtrackingSolver(visited, neighbor, end); ok {
				neighbor.isPath = true
				return true
			}
		}
	}

	return false
}

func (m *Maze) DijkstraSolver() {

}
