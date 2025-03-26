package report

import (
	"codeberg.org/go-pdf/fpdf"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	font := Font{
		Name:  "Arial",
		Style: "",
		Size:  18,
		Color: Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
	}

	lb := NewLogicalBox()

	tb1 := NewTitledBox(
		TitledBoxOptions{
			Title: "Hello world 1",
			Font:  font,
		},
		Position{
			Offset: Offset{5, 5},
			Size:   Size{0, 100},
		},
	)

	totalHeight := 0.0
	for i := range 4 {

		tb := NewTitledBox(
			TitledBoxOptions{
				Title: fmt.Sprintf("Hello world %d", i),
				Font:  font,
			},
			Position{
				Offset: Offset{5, 5 + float64(i*15)},
				Size:   Size{100, 10},
			},
		)

		err := tb1.AddElement(tb)
		if err != nil {
			t.Error(tb)
		}

		totalHeight += 15
	}

	tb1.SetSize(120, totalHeight+5)

	err := lb.AddElement(tb1)
	if err != nil {
		t.Error(err)
	}

	doc := fpdf.New("P", "mm", "A4", "")

	doc.AddPage()

	err = lb.Render(doc, Position{
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
