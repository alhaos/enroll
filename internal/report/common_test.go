package report

import (
	"codeberg.org/go-pdf/fpdf"
	"fmt"
	"testing"
)

func TestCommon(t *testing.T) {

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

	b := NewBox(
		Position{
			Offset: Offset{10, 10},
			Size:   Size{277, 250},
		},
		"F",
		Color{
			200,
			200,
			200,
		},
	)

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
	for i := range 8 {

		tb := NewTitledBox(
			TitledBoxOptions{
				Title: fmt.Sprintf("Hello world %d", i),
				Font:  font,
			},
			Position{
				Offset: Offset{5, 5 + float64(i*12)},
				Size:   Size{100, 10},
			},
		)

		err := tb1.AddElement(tb)
		if err != nil {
			t.Error(tb)
		}

		totalHeight += 12
	}

	tb1.SetSize(120, totalHeight+5)

	err := b.AddElement(tb1)
	if err != nil {
		t.Error(err)
	}

	doc := fpdf.New("P", "mm", "A4", "")

	doc.AddPage()

	err = b.Render(
		doc,
		Position{
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

func TestNewImage(t *testing.T) {

	font := Font{
		Name:  "Arial",
		Style: "",
		Size:  8,
		Color: Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		},
	}

	fillColor := Color{
		Red:   66,
		Green: 95,
		Blue:  168,
	}

	lb := NewBox(
		Position{
			Offset: Offset{0, 0},
			Size:   Size{210, 997},
		},
		"F",
		Color{
			200,
			200,
			200,
		},
	)

	tb := NewTitledBox(
		TitledBoxOptions{
			Title:     "Heloo world",
			Font:      font,
			FillColor: fillColor,
		},
		Position{
			Offset: Offset{5, 5},
			Size:   Size{200, 190},
		},
	)

	err := lb.AddElement(tb)
	if err != nil {
		t.Error(err)
	}

	img := NewImage(
		`D:\repo\enroll\resourcees\image\Sample-PNG-Image.png`,
		Position{
			Offset: Offset{5, 5},
			Size:   Size{50, 50},
		},
	)

	img2 := NewImage(
		`D:\repo\enroll\resourcees\image\Sample-PNG-Image.png`,
		Position{
			Offset: Offset{60, 5},
			Size:   Size{50, 50},
		},
	)

	err = tb.AddElement(img)
	if err != nil {
		t.Error(err)
	}

	err = tb.AddElement(img2)
	if err != nil {
		t.Error(err)
	}

	doc := fpdf.New(
		fpdf.OrientationPortrait,
		fpdf.UnitMillimeter,
		fpdf.PageSizeA4,
		"",
	)

	doc.AddPage()

	err = lb.Render(
		doc,
		Position{
			Offset: Offset{0, 0},
			Size:   Size{210, 297},
		},
	)
	if err != nil {
		t.Error(err)
	}

	err = doc.OutputFileAndClose("testWithImage.pdf")
	if err != nil {
		t.Error(err)
	}

}

func TestNewLabel(t *testing.T) {

	font := Font{
		Name:  "Arial",
		Style: "",
		Size:  8,
		Color: Color{
			Red:   255,
			Green: 255,
			Blue:  255,
		},
	}

	fillColor := Color{
		Red:   66,
		Green: 95,
		Blue:  168,
	}

	b := NewBox(
		Position{
			Offset: Offset{0, 0},
			Size:   Size{210, 997},
		},
		"F",
		Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		},
	)

	tb := NewTitledBox(TitledBoxOptions{
		Title:     "Titled box header",
		Font:      font,
		FillColor: fillColor,
	}, Position{
		Offset: Offset{5, 5},
		Size:   Size{200, 190},
	})

	err := b.AddElement(tb)
	if err != nil {
		t.Error(err)
	}

	font.Color = Color{
		Red:   0,
		Green: 0,
		Blue:  0,
	}

	for i := range 3 {
		l := NewLabel(font, "Label text", Position{
			Offset: Offset{
				X: 2,
				Y: 2 + float64(i)*(font.Size*pixelToMM+1),
			},
			Size: Size{
				Width:  20,
				Height: 20,
			},
		})
		err = tb.AddElement(l)
		if err != nil {
			t.Error(err)
		}

	}

	doc := fpdf.New(
		fpdf.OrientationPortrait,
		fpdf.UnitMillimeter,
		fpdf.PageSizeA4,
		"",
	)

	doc.AddPage()

	err = b.Render(doc, Position{
		Offset: Offset{0, 0},
		Size:   Size{210, 297},
	})
	if err != nil {
		t.Error(err)
	}

	err = doc.OutputFileAndClose("testWithLabel.pdf")
	if err != nil {
		t.Error(err)
	}

}
