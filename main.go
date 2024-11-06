package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func parseKeyStrokes(stroke string) string {
	r := regexp.MustCompile(`-?\d+(\.\d+)?`)
	allNums := r.FindAllString(stroke, 10)
	fmt.Printf("This is my array: %v\n\n", allNums)

	asciCode, err := strconv.Atoi(allNums[1])

	if err != nil {
		return "CANT DO IT"
	}

	return fromKeyCodeToString(asciCode)
}

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}

	wid, _ := xproto.NewWindowId(X)
	screen := xproto.Setup(X).DefaultScreen(X)
	context, err := xproto.NewGcontextId(X)
	xproto.CreateGC(
		X,
		context,
		xproto.Drawable(screen.Root),
		xproto.GcForeground|xproto.GcGraphicsExposures,
		[]uint32{screen.BlackPixel, 0},
	)
	xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 10,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			screen.WhitePixel,
			xproto.EventMaskStructureNotify |
				xproto.EventMaskExposure |
				xproto.EventMaskKeyRelease})

	xproto.MapWindow(X, wid)

	for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			// fmt.Printf("Event: %s\n", ev)
			fmt.Println("Drawing")
			// fmt.Printf("Event: %s\n", parseKeyStrokes(ev.String()))
			// key := parseKeyStrokes(ev.String())
			key := "1"
			var x, y int16
			x += 10
			y += 10
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
		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}

func fromKeyCodeToString(code int) string {
	switch code {
	case 10:
		return "1"
	case 11:
		return "2"
	case 12:
		return "3"
	case 13:
		return "4"
	case 14:
		return "5"
	case 15:
		return "6"
	case 16:
		return "7"
	case 17:
		return "8"
	case 18:
		return "9"
	case 19:
		return "0"

	default:
		return ""
	}
}
