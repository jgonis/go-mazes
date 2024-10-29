package grid

import (
	"fmt"
	"image"
	"testing"
)

func TestGridString(t *testing.T) {
	grid := NewGrid(3, 3)
	want :=
		`+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+`
	got := grid.String()
	if want != got {
		t.Errorf("want: \n%s\n got: \n%s\n", want, got)
	}

	grid.CellAt(image.Point{0, 0}).Link(*grid.CellAt(image.Point{1, 0}), true)
	grid.CellAt(image.Point{1, 1}).Link(*grid.CellAt(image.Point{2, 1}), true)
	grid.CellAt(image.Point{1, 1}).Link(*grid.CellAt(image.Point{1, 0}), true)
	got = grid.String()
	fmt.Println(got)
}
