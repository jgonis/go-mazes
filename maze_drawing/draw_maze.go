package maze_drawing

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/jgonis/go-mazes/mazegrid"
)

const (
	dpi             = 300
	letterHeight    = 11
	letterWidth     = 8.5
	letterHeightPix = letterHeight * dpi
	letterWidthPix  = letterWidth * dpi
	letterBorderPix = 0.5 * dpi
)

type PaperOrientation int

const (
	Portrait PaperOrientation = iota
	Landscape
)

func DrawMaze(grid *mazegrid.Grid) string {
	var orientation PaperOrientation
	var cellSize uint
	if grid.Width >= grid.Height {
		orientation = Landscape
		cellSize = (letterWidthPix - (2 * letterBorderPix)) / grid.Width
	} else {
		orientation = Portrait
		cellSize = (letterHeightPix - (2 * letterBorderPix)) / grid.Height
	}
	drawingContext := createDrawingContext(orientation)
	drawingContext.SetRGB(0, 0, 0)
	drawingContext.SetLineWidth(8)
	// iterate across each cell in the grid
	for row := range grid.Height {
		for column := range grid.Width {
			cell := grid.CellAt(image.Point{int(column), int(row)})
			// calculate a new center point based on the offset of the cell in the grid
			cellCenterX, cellCenterY := calculateCenterPoint(row, column, cellSize)

			// for each of the cell's neighbors
			// if the neighbor is nil or the cell is not linked to it, draw a line
			if cell.North == nil || !cell.IsCellLinked(*cell.North) {
				x1 := float64(cellCenterX - (cellSize / 2))
				y1 := float64(cellCenterY - (cellSize / 2))
				x2 := float64(cellCenterX + (cellSize / 2))
				y2 := float64(cellCenterY - (cellSize / 2))
				drawingContext.DrawLine(x1, y1, x2, y2)
			}
			if cell.East == nil || !cell.IsCellLinked(*cell.East) {
				x1 := float64(cellCenterX + (cellSize / 2))
				y1 := float64(cellCenterY - (cellSize / 2))
				x2 := float64(cellCenterX + (cellSize / 2))
				y2 := float64(cellCenterY + (cellSize / 2))
				drawingContext.DrawLine(x1, y1, x2, y2)
			}
			if cell.South == nil || !cell.IsCellLinked(*cell.South) {
				x1 := float64(cellCenterX - (cellSize / 2))
				y1 := float64(cellCenterY + (cellSize / 2))
				x2 := float64(cellCenterX + (cellSize / 2))
				y2 := float64(cellCenterY + (cellSize / 2))
				drawingContext.DrawLine(x1, y1, x2, y2)
			}
			if cell.West == nil || !cell.IsCellLinked(*cell.West) {
				x1 := float64(cellCenterX - (cellSize / 2))
				y1 := float64(cellCenterY - (cellSize / 2))
				x2 := float64(cellCenterX - (cellSize / 2))
				y2 := float64(cellCenterY + (cellSize / 2))
				drawingContext.DrawLine(x1, y1, x2, y2)
			}
			drawingContext.Stroke()
		}
	}
	drawingContext.SavePNG("maze.png")
	return "Drawing saved to test.png"
}

func calculateCenterPoint(cellRow, cellColumn, cellSize uint) (uint, uint) {
	cellCenterX := letterBorderPix + (cellSize * cellRow) + (cellSize / 2)
	cellCenterY := letterBorderPix + (cellSize * cellColumn) + (cellSize / 2)

	return cellCenterX, cellCenterY
}

func createDrawingContext(orientation PaperOrientation) *gg.Context {
	var dc *gg.Context
	if orientation == Landscape {
		dc = gg.NewContext(letterHeightPix, letterWidthPix)

	} else {
		dc = gg.NewContext(letterWidthPix, letterHeightPix)
	}
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	return dc
}
