package main

import (
	"flag"

	"github.com/jgonis/go-mazes/maze_algorithms"
	"github.com/jgonis/go-mazes/maze_drawing"
	"github.com/jgonis/go-mazes/mazegrid"
)

func main() {
	widthArg := flag.Uint("width", 4, "width of the maze")
	heightArg := flag.Uint("height", 3, "height of the maze")
	algoArg := flag.String("algorithm", "binary_tree", "algorithm to use to generate the maze")
	flag.Parse()

	if *widthArg == 0 || *heightArg == 0 {
		panic("width and height must be greater than 0")
	}

	grid := mazegrid.NewGrid(*widthArg, *heightArg)
	newGrid := maze_algorithms.GetMazeGenerator(*algoArg).GenerateMaze(&grid)
	newGrid.CreateStartAndEnd()

	maze_drawing.DrawMaze(newGrid)
}
