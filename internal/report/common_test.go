package report

import (
	"codeberg.org/go-pdf/fpdf"
	"testing"
)

func TestNew(t *testing.T) {

	font := Font{
		Name:  "Arial",
		Style: "",
		Size:  8,
		Color: Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
	}

	report := NewLogicalBox()

	tb1 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 1",
			Font:  font,
		},
		Position{
			Offset: Offset{5, 5},
			Size:   Size{195, 100},
		},
	)

	err := report.AddElement(tb1)
	if err != nil {
		t.Error(err)
	}

	tb2 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 2",
			Font:  font,
		},
		Position{
			Offset: Offset{5, 5},
			Size:   Size{90, 90},
		},
	)

	err = tb1.AddElement(tb2)
	if err != nil {
		t.Error(err)
	}

	tb3 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 3",
			Font:  font,
		},
		Position{
			Offset: Offset{100, 5},
			Size:   Size{90, 90},
		},
	)

	err = tb1.AddElement(tb3)
	if err != nil {
		t.Error(err)
	}

	tb4 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 4",
			Font:  font,
		},
		Position{
			Offset: Offset{5, 5},
			Size:   Size{50, 50},
		},
	)

	err = tb2.AddElement(tb4)
	if err != nil {
		t.Error(err)
	}

	tb5 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 5",
			Font:  font,
		},
		Position{
			Offset: Offset{5, 60},
			Size:   Size{50, 25},
		},
	)

	err = tb2.AddElement(tb5)
	if err != nil {
		t.Error(err)
	}

	doc := fpdf.New("P", "mm", "A4", "")
	doc.AddPage()

	err = report.Render(doc, Position{
		Offset: Offset{0, 0},
		Size:   Size{300, 300},
	})
	if err != nil {
		t.Error(err)
	}

	err = doc.OutputFileAndClose("test.pdf")
	if err != nil {
		t.Error(err)
	}
}
