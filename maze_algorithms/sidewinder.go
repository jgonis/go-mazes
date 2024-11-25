package maze_algorithms

import (
	"image"

	"math/rand"

	"github.com/jgonis/go-mazes/mazegrid"
)

type Direction uint

const (
	East Direction = iota
	North
)

type Sidewinder struct{}

func (s Sidewinder) GenerateMaze(grid *mazegrid.Grid) *mazegrid.Grid {
	// start in bottom west corner
	for i := int(grid.Height - 1); i >= 0; i-- {
		run := []*mazegrid.Cell{}
		for j := range grid.Width {
			cell := grid.CellAt(image.Point{int(j), int(i)})
			// add current cell to the current run of cells
			run = append(run, cell)

			// flip a coin to decide if we carve east or north
			choice := Direction(rand.Intn(2))

			// if the coin flip says east, and we can carve east (ie we're not on the edge)
			if choice == East {
				// if we can't carve east
				if cell.East == nil {
					// then just try to carve north
					if cell.North != nil {
						cell.Link(cell.North, true)
						clear(run)
						run = []*mazegrid.Cell{}
					}
					// if we can't carve north (ie we are at the top) then just continue on
				} else {
					cell.Link(cell.East, true)
				}
				// if the coin flip say carve north
			} else {
				// and we can carve north,
				if cell.North != nil {
					// randomly choose one
					index := rand.Intn(len(run))
					// of the cells in the run to carve north with
					run[index].Link(run[index].North, true)
					// if we can't then try to carve east
				} else {
					if cell.East != nil {
						cell.Link(cell.East, true)

					}
					// if we can't carve east then continue on
				}
				// whichever we chose for the carve-north case, clear the current run of cells
				run = []*mazegrid.Cell{}
			}

		}
	}
	return grid
}
