package maze_algorithms

import (
	"image"
	"math/rand"

	"github.com/jgonis/go-mazes/mazegrid"
)

type Sidewinder struct{}

func (s Sidewinder) GenerateMaze(grid *mazegrid.Grid) *mazegrid.Grid {
	for i := range grid.Height {
		for j := range grid.Width {
			cell := grid.CellAt(image.Point{int(j), int(i)})
			neighbors := []*mazegrid.Cell{}
			if cell.North != nil {
				neighbors = append(neighbors, cell.North)
			}
			if cell.East != nil {
				neighbors = append(neighbors, cell.East)
			}

			if len(neighbors) != 0 {
				index := rand.Intn(len(neighbors))
				neighbor := neighbors[index]
				cell.Link(neighbor, true)
			}
		}
	}
	return grid
}
