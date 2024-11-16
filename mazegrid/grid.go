package mazegrid

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

func (g Grid) String() string {
	mazeString := ""
	for rowIndex := range g.Height {
		row := g.Cells[int(rowIndex)]
		//draw top of the row
		for _, cell := range row {
			if cell.North == nil {
				mazeString += "+---"
			} else if cell.North != nil {
				if cell.IsCellLinked(*cell.North) {
					mazeString += "+   "
				} else {
					mazeString += "+---"
				}
			}
		}
		mazeString += "+\n"
		//draw walls of the row
		for _, cell := range row {
			if cell.West == nil {
				mazeString += "|   "
			} else if cell.West != nil {
				if cell.IsCellLinked(*cell.West) {
					mazeString += "    "
				} else {
					mazeString += "|   "
				}
			}
		}
		mazeString += "|\n"
	}
	//draw bottom of the grid
	for range g.Width {
		mazeString += "+---"
	}
	//append final + to the string
	mazeString += "+"
	return mazeString
}

func initializeGridCells(width, height uint, grid *Grid, cells map[int][]Cell) map[int][]Cell {
	for row := range height {
		for column := range width {
			cell := grid.CellAt(image.Point{int(column), int(row)})
			cell.North = grid.CellAt(image.Point{int(column), int(row) - 1})
			cell.South = grid.CellAt(image.Point{int(column), int(row) + 1})
			cell.East = grid.CellAt(image.Point{int(column) + 1, int(row)})
			cell.West = grid.CellAt(image.Point{int(column) - 1, int(row)})
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
