package main

import (
	// #include "draw.h"
	// #cgo LDFLAGS: -L/usr/X11/lib -lX11 -lstdc++
	"C"
)

func main() {
	C.draw()
}

