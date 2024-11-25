package maze_algorithms

import "github.com/jgonis/go-mazes/mazegrid"

type MazeGenerator interface {
	GenerateMaze(grid *mazegrid.Grid) *mazegrid.Grid
}

func GetMazeGenerator(algo string) MazeGenerator {
	switch algo {
	case "binary_tree":
		return &BinaryTree{}
	case "sidewinder":
		return &Sidewinder{}
	default:
		panic("invalid algorithm")
	}
}
