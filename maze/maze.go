package maze

import (
	"math/rand"
)

type Maze struct {
	width         int
	height        int
	wallThickness int
	cellThickness int
	cells         [][]*Cell
}

func NewMaze(width, height, wallThickness, cellThickness int) *Maze {
	rand.Seed(int64(width * height))

	m := &Maze{
		width:         width,
		height:        height,
		wallThickness: wallThickness,
		cellThickness: cellThickness,
		cells:         make([][]*Cell, height),
	}

	for row := 0; row < height; row++ {
		var rowCells []*Cell
		for column := 0; column < width; column++ {
			rowCells = append(rowCells, &Cell{})
		}
		m.cells[row] = rowCells
	}

	for row := 0; row < m.height; row++ {
		for column := 0; column < m.width; column++ {
			c, right, bottom := m.cell(column, row), m.cell(column+1, row), m.cell(column, row+1)
			c.connect(Right, right)
			c.connect(Bottom, bottom)
		}
	}

	m.Reset()

	return m
}

func (m *Maze) Reset() {
	for row := 0; row < m.height; row++ {
		for column := 0; column < m.width; column++ {
			m.cell(row, column).isPath = false
		}
	}
}

func (m *Maze) cell(column, row int) *Cell {
	if column < 0 || m.width <= column ||
		row < 0 || m.height <= row {
		return nil
	}
	return m.cells[row][column]
}
