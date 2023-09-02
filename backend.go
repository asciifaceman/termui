// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package tooey

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var scrn tcell.Screen

// Init is refactor of init for tcell operation
func Init() error {
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err = s.Init(); err != nil {
		return err
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	s.SetStyle(defaultStyle)
	s.Clear()

	scrn = s

	return nil
}

// Close is refactor of close
func Close() {
	maybePanic := recover()
	scrn.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

// DrawableDimensions is the same as TerminalDimensions -1 to represent visibly drawable space in
// most terminals
func DrawableDimensions() (int, int) {
	width, height := TerminalDimensions()
	return width - 1, height - 1
}

func TerminalDimensions() (int, int) {
	scrn.Sync()
	width, height := scrn.Size()
	return width, height
}

func Clear() {
	scrn.Clear()
	//tb.Clear(tb.ColorDefault, tb.Attribute(Theme.Default.Bg+1))
}
