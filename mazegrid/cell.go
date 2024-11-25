package mazegrid

import "image"

type Cell struct {
	coordinates image.Point
	linkedCells map[*Cell]bool
	IsStart     bool
	IsEnd       bool
	grid        *Grid
	North       *Cell
	South       *Cell
	East        *Cell
	West        *Cell
}

func NewCell(x, y uint, grid *Grid) Cell {
	return Cell{
		coordinates: image.Point{int(x), int(y)},
		linkedCells: make(map[*Cell]bool),
		grid:        grid,
	}
}

func (c *Cell) Link(cell *Cell, bidirectional bool) {
	c.linkedCells[cell] = true
	if bidirectional {
		cell.Link(c, false)
	}
}

func (c *Cell) Unlink(cell *Cell, bidirectional bool) {
	c.linkedCells[cell] = false
	if bidirectional {
		cell.Unlink(c, false)
	}
}

func (c Cell) GetLinkedCells() []*Cell {
	var linkedCells []*Cell
	for cell, present := range c.linkedCells {
		if present {
			linkedCells = append(linkedCells, cell)
		}
	}
	return linkedCells
}

func (c Cell) IsCellLinked(cell *Cell) bool {
	_, present := c.linkedCells[cell]
	return present
}

func (c Cell) Neighbors() []Cell {
	neighbors := []Cell{}
	if c.North != nil {
		neighbors = append(neighbors, *c.North)
	}
	if c.South != nil {
		neighbors = append(neighbors, *c.South)
	}
	if c.East != nil {
		neighbors = append(neighbors, *c.East)
	}
	if c.West != nil {
		neighbors = append(neighbors, *c.West)
	}
	return neighbors
}
