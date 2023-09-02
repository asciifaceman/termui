# Tooey

This is a FORK of gizak/termui see [FORK.MD](FORK.md) for details.

[<img src="./_assets/demo.gif" alt="demo cast under osx 10.10; Terminal.app; Menlo Regular 12pt.)" width="100%">](./_examples/demo.go)

tooey is a cross-platform and customizable terminal dashboard and widget library built on top of [tcell](https://github.com/gdamore/tcell) (formerly termbox-go pre-fork). This project was inspired by its parent, gizak/termui.

## Features

- 24bit colors enabled by TCell
	- Requires your $TERM to end in `-truecolor`
	- ex. `export TERM=screen-truecolor`
- Several premade widgets for common use cases
- Easily create custom widgets
- Position widgets either in a relative grid or with absolute coordinates
- Keyboard, mouse, and terminal resizing events
- Colors and styling

## Installation

`go get github.com/asciifaceman/tooey`

## Hello World

```go
package main

import (
	"log"

	ui "github.com/asciifaceman/tooey""
	"github.com/asciifaceman/tooey/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
```

## Widgets

- [BarChart](./_examples/barchart.go)
- [Canvas](./_examples/canvas.go) (for drawing braille dots)
- [Gauge](./_examples/gauge.go)
- [Image](./_examples/image.go)
- [List](./_examples/list.go)
- [Tree](./_examples/tree.go)
- [Paragraph](./_examples/paragraph.go)
- [PieChart](./_examples/piechart.go)
- [Plot](./_examples/plot.go) (for scatterplots and linecharts)
- [Sparkline](./_examples/sparkline.go)
- [StackedBarChart](./_examples/stacked_barchart.go)
- [Table](./_examples/table.go)
- [Tabs](./_examples/tabs.go)

Run an example with `go run _examples/{example}.go` or run each example consecutively with `make run-examples`.

## Documentation

- [wiki](https://github.com/asciifaceman/tooey/wiki)

## License

[MIT](http://opensource.org/licenses/MIT)
