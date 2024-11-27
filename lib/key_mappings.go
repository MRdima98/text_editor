package lib

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseKeyStrokes(stroke string, uppercase bool) string {
	r := regexp.MustCompile(`-?\d+(\.\d+)?`)
	allNums := r.FindAllString(stroke, 10)

	asciCode, err := strconv.Atoi(allNums[1])

	if err != nil {
		return "CANT DO IT"
	}

	return fromKeyCodeToString(asciCode, uppercase)
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
	case 23:
		return TAB
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
