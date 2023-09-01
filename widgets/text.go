// TCell implementation of paragraph
//
// Charles <asciifaceman> Corbett 2023
//
//

package widgets

import (
	"github.com/asciifaceman/tooey"
	"github.com/gdamore/tcell"
)

// Text represents a simple styled block of text with wrapping
// paragraph capabilities
type Text struct {
	tooey.Element
	Content   string
	TextStyle tooey.Style // TODO: Style type
	Wrap      bool
}

// NewText returns a basic empty *Text
func NewText() *Text {
	return &Text{
		Wrap:      true,
		TextStyle: tooey.StyleDefault,
	}
}

// Draw ...
func (t *Text) Draw(s tcell.Screen) {
	t.Element.Draw(s)

	row := t.Rectangle.Min.Y
	col := t.Rectangle.Min.X

	for _, r := range t.Content {

		// TODO: Handle zero width characters

		s.SetContent(col, row, r, nil, t.TextStyle.Style)
		col++

		if t.Wrap {
			if col > t.Rectangle.Max.X {
				row++
				col = t.Rectangle.Min.X
			}
			if row > t.Rectangle.Max.Y {
				// gobble the remainder
				break
			}
		}
	}

}
