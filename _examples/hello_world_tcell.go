//go:build ignore
// +build ignore

package main

import (
	"log"
	"time"

	. "github.com/asciifaceman/tooey"
	widgets "github.com/asciifaceman/tooey/widgets"
)

func main() {
	if err := Init(); err != nil {
		log.Fatalf("failed to initialize tooey: %v", err)
	}
	defer Close()

	cont := widgets.NewContainer()
	cont.Title = "This Container"
	cont.SetRect(5, 5, 35, 15)

	// make simple paragraph in box
	t := widgets.NewText()
	t.Content = "Testing some longer strings that should wrap a smaller geometry"
	t.Wrap = true
	t.SetRect(1, 1, 15, 15)

	cont.Append(t)

	Render(cont)

	time.Sleep(time.Duration(time.Second * 5))

}
