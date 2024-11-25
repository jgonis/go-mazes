package mazegrid

import (
	"image"
	"testing"
)

func TestUnlinkedGridString(t *testing.T) {
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
}

func TestLinkedGridString(t *testing.T) {
	grid := NewGrid(3, 3)

	grid.CellAt(image.Point{0, 0}).Link(grid.CellAt(image.Point{1, 0}), true)
	grid.CellAt(image.Point{1, 1}).Link(grid.CellAt(image.Point{2, 1}), true)
	grid.CellAt(image.Point{1, 1}).Link(grid.CellAt(image.Point{1, 0}), true)
	got := grid.String()
	want :=
		`+---+---+---+
|       |   |
+---+   +---+
|   |       |
+---+---+---+
|   |   |   |
+---+---+---+`
	if want != got {
		t.Errorf("want: \n%s\n got: \n%s\n", want, got)
	}
}
