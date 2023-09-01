package tooey

import (
	"image"
	"sync"

	"github.com/gdamore/tcell"
)

// Element is the base struct inherited by most widgets
// It is a port of the original Block
// Element manages size and position.
// It does nothing else. All other widgets and ui elements will
// oerride the Draw method
type Element struct {
	PaddingLeft   int
	PaddingRight  int
	PaddingTop    int
	PaddingBottom int

	image.Rectangle
	Inner image.Rectangle

	sync.Mutex
}

// Draw implements the Drawable interface
func (e *Element) Draw(s tcell.Screen) {
}

// SetRect implements the Drawable interface
func (e *Element) SetRect(x1 int, y1 int, x2 int, y2 int) {
	e.Rectangle = image.Rect(x1, y1, x2, y2)
	e.Inner = image.Rect(
		e.Min.X+e.PaddingLeft,
		e.Min.Y+e.PaddingTop,
		e.Max.X-e.PaddingRight,
		e.Max.Y-e.PaddingBottom,
	)
}

// X1 returns the rects Min.X
func (e *Element) X1() int {
	return e.Rectangle.Min.X
}

// X2 returns the rects Max.X
func (e *Element) X2() int {
	return e.Rectangle.Max.X
}

// Y1 returns the rects Min.X
func (e *Element) Y1() int {
	return e.Rectangle.Min.Y
}

// Y2 returns the rects Min.X
func (e *Element) Y2() int {
	return e.Rectangle.Max.Y
}

// GetRect implements the Drawable interface
func (e *Element) GetRect() image.Rectangle {
	return e.Rectangle
}

// Container is an element that can hold other elements within
// this needs a lot more thought out
// in the context of tcell
type Container struct {
	Element
	Contents []Drawable
}

func (c *Container) Draw(s tcell.Screen) {
	for _, drawable := range c.Contents {
		drawable.Draw(s)
	}
}
