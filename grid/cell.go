package grid

import "image"

type Cell struct {
	coordinates image.Point
	linkedCells map[image.Point]bool
	grid        *Grid
	North       *Cell
	South       *Cell
	East        *Cell
	West        *Cell
}

func NewCell(x, y uint, grid *Grid) Cell {
	return Cell{
		coordinates: image.Point{int(x), int(y)},
		linkedCells: make(map[image.Point]bool),
		grid:        grid,
	}
}

func (c *Cell) Link(cell Cell, bidirectional bool) {
	c.linkedCells[cell.coordinates] = true
	if bidirectional {
		cell.Link(*c, false)
	}
}

func (c *Cell) Unlink(cell Cell, bidirectional bool) {
	c.linkedCells[cell.coordinates] = false
	if bidirectional {
		cell.Unlink(*c, false)
	}
}

func (c Cell) GetLinkedCells() []*Cell {
	var linkedCells []*Cell
	for coords := range c.linkedCells {
		linkedCells = append(linkedCells, c.grid.CellAt(coords))
	}
	return linkedCells
}

func (c Cell) IsCellLinked(cell Cell) bool {
	_, present := c.linkedCells[cell.coordinates]
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
