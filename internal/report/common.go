package report

import (
	"codeberg.org/go-pdf/fpdf"
)

const pixelToMM = 0.3528

// Element interface
type Element interface {
	AddElement(Element) error
	Render(doc *fpdf.Fpdf, position Position) error
	ChildElements() []Element
	SetSize(width float64, height float64)
}

// Font this structure contains information about font
type Font struct {
	Name  string
	Style string
	Size  float64
	Color Color
}

// Color struct
type Color struct {
	Red   int
	Green int
	Blue  int
}

// Position ...
type Position struct {
	Offset
	Size
}

// Offset contains the coordinates of the offset relative to the beginning of the document
type Offset struct {
	X float64
	Y float64
}

// Size contains the size of rectangle object
type Size struct {
	Width  float64
	Height float64
}

// NewLogicalBox create new logical box
func NewLogicalBox() *LogicalBox {
	return &LogicalBox{}
}

// LogicalBox struct
type LogicalBox struct {
	RectangleElement
}

// AddElement add child element to logical box
func (br *LogicalBox) AddElement(element Element) error {
	br.Elements = append(br.Elements, element)
	return nil
}

// ChildElements return child elements
func (br *LogicalBox) ChildElements() []Element {
	return br.Elements
}

// TitledBoxOptions ...
type TitledBoxOptions struct {
	Title string
	Font  Font
}

// TitledBox struct
type TitledBox struct {
	RectangleElement
	Options TitledBoxOptions
}

// NewTitledBox constructor for title box object
func NewTitledBox(titledBoxOptions TitledBoxOptions, position Position) *TitledBox {
	return &TitledBox{
		Options: titledBoxOptions,
		RectangleElement: RectangleElement{
			X:      position.X,
			Y:      position.Y,
			Width:  position.Width,
			Height: position.Height,
		},
	}
}

// Render titled box
func (tb *TitledBox) Render(doc *fpdf.Fpdf, position Position) error {

	const titleMargin = 1.0

	doc.SetFont(
		tb.Options.Font.Name,
		tb.Options.Font.Style,
		tb.Options.Font.Size,
	)

	doc.Text(
		position.Offset.X+tb.RectangleElement.X+titleMargin,
		position.Offset.Y+tb.Y+tb.Options.Font.Size*pixelToMM,
		tb.Options.Title,
	)

	doc.Line(
		position.Offset.X+tb.X,
		position.Offset.Y+tb.Y+tb.Options.Font.Size*pixelToMM+titleMargin,
		position.Offset.X+tb.X+tb.Width,
		position.Offset.Y+tb.Y+tb.Options.Font.Size*pixelToMM+titleMargin,
	)

	err := tb.RectangleElement.Render(doc, position)
	if err != nil {
		return err
	}
	return nil
}

// AddElement add child element to titled box
func (tb *TitledBox) AddElement(element Element) error {
	tb.Elements = append(tb.Elements, element)
	return nil
}

// ChildElements return text box child elements
func (tb *TitledBox) ChildElements() []Element {
	return tb.Elements
}

// RectangleElement composite structure for any rectangle element for
type RectangleElement struct {
	X, Y, Width, Height float64
	Elements            []Element
}

// SetSize set rectangle element size
func (re *RectangleElement) SetSize(width float64, height float64) {
	re.Width = width
	re.Height = height
}

// Render Rectangle
func (re *RectangleElement) Render(doc *fpdf.Fpdf, position Position) error {

	position.X += re.X
	position.Y += re.Y

	doc.Rect(position.X, position.Y, re.Width, re.Height, "D")

	for _, element := range re.Elements {
		err := element.Render(doc, position)
		if err != nil {
			return err
		}
	}
	return nil

}
