package main

import (
	"github.com/jgonis/go-mazes/maze_algorithms"
	"github.com/jgonis/go-mazes/maze_drawing"
	"github.com/jgonis/go-mazes/mazegrid"
)

func main() {
	grid := mazegrid.NewGrid(1000, 1000)
	newGrid := maze_algorithms.GenerateBinaryTreeMaze(&grid)
	// fmt.Println(newGrid.String())
	maze_drawing.DrawMaze(newGrid)
}
