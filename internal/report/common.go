package report

import (
	"codeberg.org/go-pdf/fpdf"
	"errors"
	"fmt"
	"os"
)

const pixelToMM = 0.3528

// Element interface
type Element interface {
	AddElement(Element) error
	Render(doc *fpdf.Fpdf, parentPosition Position) error
	ChildElements() []Element
	SetSize(width float64, height float64)
}

// RectangleElement composite structure for any rectangle element for
type RectangleElement struct {
	Position Position
	Elements []Element
}

// SetSize set rectangle element size
func (re *RectangleElement) SetSize(width float64, height float64) {
	re.Position.Width = width
	re.Position.Height = height
}

// Render Rectangle
func (re *RectangleElement) Render(doc *fpdf.Fpdf, patentPosition Position) error {

	patentPosition.X += re.Position.X
	patentPosition.Y += re.Position.Y

	doc.Rect(patentPosition.X, patentPosition.Y, re.Position.Width, re.Position.Height, "D")

	for _, element := range re.Elements {
		err := element.Render(doc, patentPosition)
		if err != nil {
			return err
		}
	}
	return nil

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

// LogicalBox struct
type LogicalBox struct {
	RectangleElement
}

// NewLogicalBox constructor from logical box
func NewLogicalBox(position Position) *LogicalBox {
	return &LogicalBox{
		RectangleElement{
			Position: position,
		},
	}
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
			Position: position,
		},
	}
}

// Render titled box
func (tb *TitledBox) Render(doc *fpdf.Fpdf, parentPosition Position) error {

	const titleMargin = 1.0

	doc.SetFont(
		tb.Options.Font.Name,
		tb.Options.Font.Style,
		tb.Options.Font.Size,
	)

	doc.Text(
		parentPosition.Offset.X+tb.RectangleElement.Position.X+titleMargin,
		parentPosition.Offset.Y+tb.Position.Y+tb.Options.Font.Size*pixelToMM,
		tb.Options.Title,
	)

	doc.Line(
		parentPosition.Offset.X+tb.Position.X,
		parentPosition.Offset.Y+tb.Position.Y+tb.Options.Font.Size*pixelToMM+titleMargin,
		parentPosition.Offset.X+tb.Position.X+tb.Position.Width,
		parentPosition.Offset.Y+tb.Position.Y+tb.Options.Font.Size*pixelToMM+titleMargin,
	)

	err := tb.RectangleElement.Render(doc, parentPosition)
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

// Image ...
type Image struct {
	filename string
	RectangleElement
}

func NewImage(fileName string, position Position) *Image {
	return &Image{
		filename: fileName,
		RectangleElement: RectangleElement{
			Position: position,
		},
	}
}

func (i *Image) AddElement(Element) error {
	return errors.New("you can't add child element to Image")
}

func (i *Image) ChildElements() []Element {
	return nil
}

func (i *Image) Render(doc *fpdf.Fpdf, parentPosition Position) error {

	_, err := os.Stat(i.filename)
	if err != nil {
		return fmt.Errorf("image file %s not found: %w", i.filename, err)
	}

	doc.Image(
		i.filename,
		parentPosition.Offset.X+i.Position.Offset.X,
		parentPosition.Offset.Y+i.Position.Offset.Y,
		i.Position.Size.Width,
		i.Position.Size.Height,
		false, "", 0, "",
	)
	return nil
}
