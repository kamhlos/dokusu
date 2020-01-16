package main

type Position struct {
	row int
	col int
}

// Cell represents each of 81 board's cells
type Cell struct {
	Number   int
	row      int
	col      int
	invalid  bool // those have a red or green color
	active   bool
	selected bool  // selected cells have a magenta color
	marks    []int // other number candidates (if number = 0)
}

const (
	cReset      = "\x1b[0m"
	cBright     = "\x1b[1m"
	cDim        = "\x1b[2m"
	cUnderscore = "\x1b[4m"
	cBlink      = "\x1b[5m"
	cReverse    = "\x1b[7m"
	cHidden     = "\x1b[8m"

	cFgBlack   = "\x1b[30m"
	cFgRed     = "\x1b[31m"
	cFgGreen   = "\x1b[32m"
	cFgYellow  = "\x1b[33m"
	cFgBlue    = "\x1b[34m"
	cFgMagenta = "\x1b[35m"
	cFgCyan    = "\x1b[36m"
	cFgWhite   = "\x1b[37m"

	cBgBlack   = "\x1b[40m"
	cBgRed     = "\x1b[41m"
	cBgGreen   = "\x1b[42m"
	cBgYellow  = "\x1b[43m"
	cBgBlue    = "\x1b[44m"
	cBgMagenta = "\x1b[45m"
	cBgCyan    = "\x1b[46m"
	cBgWhite   = "\x1b[47m"
)
