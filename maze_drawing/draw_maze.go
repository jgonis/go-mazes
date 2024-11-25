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
	} else {
		orientation = Portrait
	}
	cellSize = calculateCellSize(grid.Width, grid.Height)
	halfCellSize := cellSize / 2
	drawingContext := createDrawingContext(orientation)
	drawingContext.SetRGB(0, 0, 0)
	drawingContext.SetLineWidth(6)
	// iterate across each cell in the grid
	for row := range grid.Height {
		for column := range grid.Width {
			cell := grid.CellAt(image.Point{int(column), int(row)})
			// calculate a new center point based on the offset of the cell in the grid
			cellCenterX, cellCenterY := calculateCellCenterPoint(grid.Width, grid.Height, row, column, cellSize)

			// for each of the cell's neighbors
			// if the neighbor is nil or the cell is not linked to it, draw a line
			if cell.North == nil || !cell.IsCellLinked(cell.North) {
				x1 := float64(cellCenterX - halfCellSize)
				y1 := float64(cellCenterY - halfCellSize)
				x2 := float64(cellCenterX + halfCellSize)
				y2 := float64(cellCenterY - halfCellSize)
				drawingContext.DrawLine(x1, y1, x2, y2)
				drawingContext.Stroke()
			}
			if (cell.East == nil && !cell.IsEnd) || (cell.East != nil && !cell.IsCellLinked(cell.East)) {
				x1 := float64(cellCenterX + halfCellSize)
				y1 := float64(cellCenterY - halfCellSize)
				x2 := float64(cellCenterX + halfCellSize)
				y2 := float64(cellCenterY + halfCellSize)
				drawingContext.DrawLine(x1, y1, x2, y2)
				drawingContext.Stroke()
			}
			if cell.South == nil || !cell.IsCellLinked(cell.South) {
				x1 := float64(cellCenterX - halfCellSize)
				y1 := float64(cellCenterY + halfCellSize)
				x2 := float64(cellCenterX + halfCellSize)
				y2 := float64(cellCenterY + halfCellSize)
				drawingContext.DrawLine(x1, y1, x2, y2)
				drawingContext.Stroke()
			}
			if (cell.West == nil && !cell.IsStart) || (cell.West != nil && !cell.IsCellLinked(cell.West)) {
				x1 := float64(cellCenterX - halfCellSize)
				y1 := float64(cellCenterY - halfCellSize)
				x2 := float64(cellCenterX - halfCellSize)
				y2 := float64(cellCenterY + halfCellSize)
				drawingContext.DrawLine(x1, y1, x2, y2)
				drawingContext.Stroke()
			}
		}
	}
	drawingContext.SavePNG("maze.png")
	return "Drawing saved to test.png"
}

func calculateCellSize(gridWidth, gridHeight uint) uint {
	var potentialCellSize1, potentialCellSize2, pageWidth, pageHeight uint
	if gridWidth >= gridHeight {
		//Landscape mode
		pageWidth = uint(letterHeightPix)
		pageHeight = uint(letterWidthPix)
		potentialCellSize1 = (pageWidth - (2 * letterBorderPix)) / gridWidth
		potentialCellSize2 = (pageHeight - (2 * letterBorderPix)) / gridHeight
	} else {
		pageWidth = uint(letterHeightPix)
		pageHeight = uint(letterWidthPix)
		potentialCellSize1 = (pageWidth - (2 * letterBorderPix)) / gridWidth
		potentialCellSize2 = (pageHeight - (2 * letterBorderPix)) / gridHeight
	}

	if potentialCellSize1*gridHeight > pageHeight || potentialCellSize1*gridWidth > pageWidth {

		if potentialCellSize2*gridHeight > pageHeight || potentialCellSize2*gridWidth > pageWidth {
			panic("can't figure out cell size!")
		} else {
			return potentialCellSize2
		}
	} else {
		return potentialCellSize1
	}
}

func calculateCellCenterPoint(gridWidth, gridHeight, cellRow, cellColumn, cellSize uint) (uint, uint) {
	cellCenterX := calculateLeftBorder(gridWidth, gridHeight, cellSize) + (cellSize * cellColumn) + (cellSize / 2)
	cellCenterY := calculateTopBorder(gridWidth, gridHeight, cellSize) + (cellSize * cellRow) + (cellSize / 2)

	return cellCenterX, cellCenterY
}

func calculateLeftBorder(gridWidth, gridHeight, cellSize uint) uint {
	// taller than we are wide
	var pageWidth uint
	if gridWidth >= gridHeight {
		pageWidth = letterHeightPix
	} else {
		pageWidth = letterWidthPix
	}
	mazeWidth := cellSize * gridWidth
	pageCenterX := uint(pageWidth / 2)
	return pageCenterX - (mazeWidth / 2)
}

func calculateTopBorder(gridWidth, gridHeight, cellSize uint) uint {
	var pageHeight uint
	if gridWidth >= gridHeight {
		pageHeight = letterWidthPix
	} else {
		pageHeight = letterHeightPix
	}
	mazeHeight := cellSize * gridHeight
	pageCenterY := uint(pageHeight / 2)
	return pageCenterY - (mazeHeight / 2)
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
