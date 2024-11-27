package main

import (
	"fmt"
	"strings"

	"editor/lib"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func main() {
	var wholeText string
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	wid, _ := xproto.NewWindowId(X)
	screen := xproto.Setup(X).DefaultScreen(X)
	context, err := xproto.NewGcontextId(X)
	if err != nil {
		fmt.Println(err)
		return
	}
	font, err := xproto.NewFontId(X)
	if err != nil {
		fmt.Println(err)
		return
	}

	xproto.OpenFont(X, font, uint16(len(lib.FONT_SIZE)), lib.FONT_SIZE)

	xproto.CreateGC(
		X,
		context,
		xproto.Drawable(screen.Root),
		xproto.GcForeground|xproto.GcBackground|xproto.GcFont,
		[]uint32{screen.WhitePixel, screen.BlackPixel, uint32(font)},
	)

	xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 10,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{
			screen.BlackPixel,
			xproto.EventMaskStructureNotify |
				xproto.EventMaskExposure |
				xproto.EventMaskKeyPress |
				xproto.EventMaskKeyRelease})

	xproto.MapWindow(X, wid)

	var x int16 = lib.STARTING_POINT
	var y int16 = lib.STARTING_POINT
	uppercase := false

	for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			fmt.Println(ev.String())
			xproto.ImageText8(
				X,
				uint8(len(lib.CURSOR)),
				xproto.Drawable(wid),
				context,
				x+lib.NEXT_LETTER,
				y,
				lib.CURSOR,
			)
		}

		if strings.Contains(ev.String(), "ConfigureNotify") || strings.Contains(ev.String(), "Expose") {
			x = lib.STARTING_POINT
			y = lib.STARTING_POINT
			for _, el := range wholeText {
				x += lib.NEXT_LETTER
				if x > lib.X_BOUND {
					x = lib.STARTING_POINT
					y += lib.NEXT_LINE
				}
				key := string(el)
				fmt.Println(key)
				xproto.ImageText8(
					X,
					uint8(len(key)),
					xproto.Drawable(wid),
					context,
					x,
					y,
					key,
				)
			}
		}

		if strings.Contains(ev.String(), "KeyPress") {
			key := lib.ParseKeyStrokes(ev.String(), uppercase)
			if key == lib.SHIFT {
				uppercase = true
			}

			if key == lib.TAB {
				for range 4 {
					x += lib.NEXT_LETTER
				}
			}

			if isCharValid(key) {
				x += lib.NEXT_LETTER
				wholeText += key
				xproto.ImageText8(
					X,
					uint8(len(key)),
					xproto.Drawable(wid),
					context,
					x,
					y,
					key,
				)

			}
		}

		if strings.Contains(ev.String(), "KeyRelease") {
			key := lib.ParseKeyStrokes(ev.String(), uppercase)
			if key == lib.SHIFT {
				uppercase = false
			}
			if x > lib.X_BOUND {
				x = lib.STARTING_POINT
				y += lib.NEXT_LINE
			}
		}

		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}

func isCharValid(key string) bool {
	return key != lib.NO_KEY && key != lib.SHIFT && key != lib.TAB
}
