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

	f := tb.Options.Font

	doc.SetFont(
		f.Name,
		f.Style,
		f.Size,
	)

	x := parentPosition.Offset.X + tb.RectangleElement.Position.X
	y := parentPosition.Offset.Y + tb.Position.Y

	doc.Text(
		x+titleMargin,
		y+f.Size*pixelToMM,
		tb.Options.Title,
	)

	doc.Line(
		x,
		y+f.Size*pixelToMM+titleMargin,
		x+tb.Position.Width,
		y+f.Size*pixelToMM+titleMargin,
	)

	doc.Rect(x, y, tb.Position.Width, tb.Position.Height, "D")

	p := Position{
		Offset: Offset{x, y + f.Size*pixelToMM + titleMargin},
		Size:   tb.Position.Size,
	}

	err := tb.RectangleElement.Render(doc, p)
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

// AddElement dummy method for implement interface element
func (i *Image) AddElement(Element) error {
	return errors.New("you can't add child element in to Image")
}

// ChildElements dummy method for implement interface element
func (i *Image) ChildElements() []Element {
	return nil
}

// Render image to doc
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

// Label element struct
type Label struct {
	RectangleElement
	Font Font
	Text string
}

// NewLabel constructor for Label element
func NewLabel(font Font, text string, position Position) *Label {
	return &Label{
		Font: font,
		Text: text,
		RectangleElement: RectangleElement{
			Position: position,
		},
	}
}

// AddElement dummy method from implement interface
func (l *Label) AddElement(Element) error {
	return errors.New("you can't add child element in to Label")
}

// ChildElements dummy element for implement interface
func (l *Label) ChildElements() []Element {
	return nil
}

func (l *Label) Render(doc *fpdf.Fpdf, parentPosition Position) error {

	f := l.Font
	doc.SetFont(f.Name, f.Style, f.Size)

	c := l.Font.Color
	doc.SetTextColor(c.Red, c.Green, c.Blue)

	doc.Text(
		parentPosition.Offset.X+l.Position.Offset.X,
		parentPosition.Offset.Y+l.Position.Offset.Y+f.Size*pixelToMM,
		l.Text,
	)
	return nil
}
