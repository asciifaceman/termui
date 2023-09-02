package tooey

import (
	"github.com/gdamore/tcell"
)

type FlexDirection uint

const (
	FlexColumn FlexDirection = iota
	FlexColumnReverse
	FlexRow
	FlexRowReverse
)

// NewContainer ...
func NewContainer() *Container {
	return &Container{
		Direction: FlexRow,
		Element:   *NewElement(),
	}
}

// Container is an element that holds other elements within
type Container struct {
	Element
	Direction FlexDirection
	Children  []ContainerChild
}

// ContainerChild which represents a member of the flex
// Grow is added up then divided by the total grow
// to find the ratio of space each child will consume
type ContainerChild struct {
	Drawable bool
	Contents interface{}
	Grow     float64
}

// NewFlexChild ...
func NewFlexChild(grow float64, i ...interface{}) ContainerChild {
	_, ok := i[0].(Drawable)
	child := i[0]
	if !ok {
		child = i
	}

	return ContainerChild{
		Drawable: ok,
		Contents: child,
		Grow:     grow,
	}
}

// Wrap embeds the given objects within the container
// using a top-level container that will fill it's available space
func (c *Container) Wrap(children ...interface{}) {
	child := ContainerChild{
		Drawable: false,
		Contents: children,
		Grow:     1.0,
	}
	c.RecursiveWrap(child)
}

// RecursiveWrap wraps a tree of children
func (c *Container) RecursiveWrap(child ContainerChild) {

	if child.Drawable {
		c.Children = append(c.Children, child)
	} else {

		children := InterfaceSlice(child.Contents)

		for i := 0; i < len(children); i++ {
			if children[i] == nil {
				continue
			}
			ch, _ := children[i].(ContainerChild)
			c.RecursiveWrap(ch)
		}

	}

}

/*
	Grow represents an amount of total Flex a child consumes

	Child 1: 1
	Child 2: 1
	Child 3: 2

	total flex: 4

	child 1 consumes: 1 /4 (0.25%)
	total width = 100 | child 1 consumes 25 width
*/

// DrawFlexRow will draw the contents as a flexible row
func (c *Container) DrawFlexRow(s tcell.Screen) {

	totalFlex := 0.0

	for _, child := range c.Children {
		totalFlex += child.Grow
	}

	width := float64(c.GetInnerRect().Dx())

	lastPosition := c.InnerX1()

	for _, child := range c.Children {
		childRatio := child.Grow / totalFlex // mult by available width
		childWidth := width * childRatio

		drawableChild := child.Contents.(Drawable)

		x := lastPosition
		y := c.InnerY1()
		w := int(childWidth)
		h := c.InnerY2()

		if x+w > c.GetInnerRect().Dx() {
			w--
		}

		drawableChild.SetRect(x, y, x+w, h)

		drawableChild.Lock()
		drawableChild.Draw(s)
		drawableChild.Unlock()

		lastPosition = x + w + 1

	}

	/*
	   //childCount := len(c.Children)
	   flexCount := 0.0

	   	for _, child := range c.Children {
	   		flexCount += child.Grow
	   	}

	   	if flexCount < 1 {
	   		flexCount = 1.0
	   	}

	   avgWidth := c.DrawableWidth() / int(flexCount)

	   	for i, child := range c.Children {
	   		if !child.Drawable {
	   			panic("somehow got non drawable child in Draw")
	   		}

	   		xratio := child.Grow / flexCount

	   		drawableChild := child.Contents.(Drawable)

	   		width := float64(c.DrawableWidth())

	   		x1 := int(width*xratio) + c.InnerX1()*i

	   		//x1 := c.InnerX1() + (avgWidth * i) + c.Padding.Left
	   		//x2 := x1 + (avgWidth * int(child.Grow)) - c.Padding.Right

	   		drawableChild.SetRect(x1, c.InnerY1(), x1+avgWidth, c.InnerY2())

	   		drawableChild.Lock()
	   		drawableChild.Draw(s)
	   		drawableChild.Unlock()

	   }
	*/
}

// DrawFlexColumn will draw the contents as a flexible column
func (c *Container) DrawFlexColumn(s tcell.Screen) {

}

//

func (c *Container) Draw(s tcell.Screen) {
	c.Element.Draw(s)

	switch c.Direction {
	case FlexColumn:
		c.DrawFlexColumn(s)
	case FlexRow:
		c.DrawFlexRow(s)
	default:
		panic("no direction set!")
	}
	/*
		childCount := len(c.Children)

		for i, child := range c.Children {
			if !child.Drawable {
				continue
			}

			avgWidth := c.DrawableWidth() / childCount

			x := c.InnerX1() + (avgWidth * i) + c.Padding.Left

			child.Contents.(Drawable).SetRect(x, c.InnerY1(), x+avgWidth-c.Padding.Right, c.InnerY2())

			child.Contents.(Drawable).Draw(s)
		}
	*/
	//
	//	childCount := len(c.Children)
	//
	//	for i, child := range c.Children {
	//		avgWidth := c.DrawableWidth() / childCount
	//		//height := c.DrawableHeight()
	//
	//		x := c.InnerX1() + (avgWidth * i) + c.Padding.Left
	//
	//		child.SetRect(x, c.InnerY1(), x+avgWidth-c.Padding.Right, c.InnerY2())
	//
	//		child.Draw(s)
	//	}

}
