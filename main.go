package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

const (
	SHIFT          = "shift"
	NO_KEY         = "noKey"
	FONT_SIZE      = "lucidasans-24"
	X_BOUND        = 450
	NEXT_LINE      = 40
	NEXT_LETTER    = 30
	STARTING_POINT = 50
)

func parseKeyStrokes(stroke string, uppercase bool) string {
	r := regexp.MustCompile(`-?\d+(\.\d+)?`)
	allNums := r.FindAllString(stroke, 10)
	// fmt.Printf("This is my array: %v\n\n", allNums)

	asciCode, err := strconv.Atoi(allNums[1])

	if err != nil {
		return "CANT DO IT"
	}

	return fromKeyCodeToString(asciCode, uppercase)
}

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

	xproto.OpenFont(X, font, uint16(len(FONT_SIZE)), FONT_SIZE)
	// _, xerr := X.WaitForEvent()
	// fmt.Println("Font errs: ", xerr)

	xproto.CreateGC(
		X,
		context,
		xproto.Drawable(screen.Root),
		xproto.GcForeground|xproto.GcBackground|xproto.GcFont,
		[]uint32{screen.WhitePixel, screen.BlackPixel, uint32(font)},
	)
	// _, xerr = X.WaitForEvent()
	// fmt.Println("GC err", xerr)

	xproto.CreateWindow(X, screen.RootDepth, wid, screen.Root,
		0, 0, 500, 500, 10,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{ // values must be in the order defined by the protocol
			screen.BlackPixel,
			xproto.EventMaskStructureNotify |
				xproto.EventMaskExposure |
				xproto.EventMaskKeyPress |
				xproto.EventMaskKeyRelease})

	xproto.MapWindow(X, wid)
	// _, xerr = X.WaitForEvent()
	// fmt.Println("Create window", xerr)

	var x int16
	var y int16 = 20
	uppercase := false

	for {
		ev, xerr := X.WaitForEvent()
		if ev == nil && xerr == nil {
			fmt.Println("Both event and error are nil. Exiting...")
			return
		}

		if ev != nil {
			fmt.Println(ev.String())
		}

		if strings.Contains(ev.String(), "ConfigureNotify") || strings.Contains(ev.String(), "Expose") {
			fmt.Println("WE ARE RESIZING")
			fmt.Println(wholeText)
			x = STARTING_POINT
			y = STARTING_POINT
			for _, el := range wholeText {
				x += NEXT_LETTER
				if x > X_BOUND {
					x = STARTING_POINT
					y += NEXT_LINE
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
			key := parseKeyStrokes(ev.String(), uppercase)
			if key == SHIFT {
				uppercase = true
			}
			if key != NO_KEY && key != SHIFT {
				x += NEXT_LETTER
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
			key := parseKeyStrokes(ev.String(), uppercase)
			if key == SHIFT {
				uppercase = false
			}
			if x > X_BOUND {
				x = STARTING_POINT
				y += NEXT_LINE
			}
		}

		if xerr != nil {
			fmt.Printf("Error: %s\n", xerr)
		}
	}
}

func fromKeyCodeToString(code int, uppercase bool) string {
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
	case 24:
		return lowerOrUpper("q", uppercase)
	case 25:
		return lowerOrUpper("w", uppercase)
	case 26:
		return lowerOrUpper("e", uppercase)
	case 27:
		return lowerOrUpper("r", uppercase)
	case 28:
		return lowerOrUpper("t", uppercase)
	case 29:
		return lowerOrUpper("y", uppercase)
	case 30:
		return lowerOrUpper("u", uppercase)
	case 31:
		return lowerOrUpper("i", uppercase)
	case 32:
		return lowerOrUpper("o", uppercase)
	case 33:
		return lowerOrUpper("p", uppercase)
	case 38:
		return lowerOrUpper("a", uppercase)
	case 39:
		return lowerOrUpper("s", uppercase)
	case 40:
		return lowerOrUpper("d", uppercase)
	case 41:
		return lowerOrUpper("f", uppercase)
	case 42:
		return lowerOrUpper("g", uppercase)
	case 43:
		return lowerOrUpper("h", uppercase)
	case 44:
		return lowerOrUpper("j", uppercase)
	case 45:
		return lowerOrUpper("k", uppercase)
	case 46:
		return lowerOrUpper("l", uppercase)
	case 50:
		return SHIFT
	case 52:
		return lowerOrUpper("z", uppercase)
	case 53:
		return lowerOrUpper("x", uppercase)
	case 54:
		return lowerOrUpper("c", uppercase)
	case 55:
		return lowerOrUpper("v", uppercase)
	case 56:
		return lowerOrUpper("b", uppercase)
	case 57:
		return lowerOrUpper("n", uppercase)
	case 58:
		return lowerOrUpper("m", uppercase)

	default:
		return NO_KEY
	}
}

func lowerOrUpper(char string, uppercase bool) string {
	if !uppercase {
		return char
	}
	return strings.ToUpper(char)
}
