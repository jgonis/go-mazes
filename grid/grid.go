package grid

import (
	"image"
	"math/rand"
)

type Grid struct {
	Width  uint
	Height uint
	Cells  map[int][]Cell
}

func NewGrid(width, height uint) Grid {
	cells := make(map[int][]Cell)
	g := Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
	cells = createGridCells(width, height, &g, cells)
	initializeGridCells(width, height, &g, cells)
	g.Cells = cells

	return g
}

func (g Grid) CellAt(coords image.Point) *Cell {
	val, present := g.Cells[coords.Y]
	if !present {
		return nil
	} else {
		if len(val) <= coords.X || coords.X < 0 {
			return nil
		}
		return &g.Cells[coords.Y][coords.X]
	}
}

func (g Grid) RandomCell() *Cell {
	randRow := rand.Intn(int(g.Height))
	randColumn := rand.Intn(int(g.Width))
	return g.CellAt(image.Point{randRow, randColumn})
}

func (g Grid) TotalCells() uint {
	return uint(g.Width * g.Height)
}

func initializeGridCells(width, height uint, grid *Grid, cells map[int][]Cell) map[int][]Cell {
	for row := range height {
		for column := range width {
			cell := grid.CellAt(image.Point{int(row), int(column)})
			cell.North = grid.CellAt(image.Point{int(row) - 1, int(column)})
			cell.South = grid.CellAt(image.Point{int(row) + 1, int(column)})
			cell.East = grid.CellAt(image.Point{int(row), int(column) + 1})
			cell.West = grid.CellAt(image.Point{int(row), int(column) - 1})
		}
	}
	return cells
}

func createGridCells(width, height uint, grid *Grid, cells map[int][]Cell) map[int][]Cell {
	for y := range height {
		cellRow := make([]Cell, 0, width)
		for x := range width {
			cellRow = append(cellRow, NewCell(x, y, grid))
		}
		cells[int(y)] = cellRow
	}
	return cells
}
