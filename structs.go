package main

type pos struct {
	row int
	col int
}

// Cell represents each of 81 board's cells
type Cell struct {
	num      int
	row      int
	col      int
	invalid  bool // those have a red or green color
	active   bool
	selected bool  // selected cells have a magenta color
	marks    []int // other num candidates (if num = 0)
}

// var board = [81]Cell{
// 	Cell{row: 0, col: 0, gen: true, num: 8},
// 	Cell{row: 0, col: 1, gen: true, num: 0},
// 	Cell{row: 0, col: 2, gen: true, num: 5},
// 	Cell{row: 0, col: 3, gen: true, num: 0},
// 	Cell{row: 0, col: 4, gen: true, num: 0},
// 	Cell{row: 0, col: 5, gen: true, num: 9},
// 	Cell{row: 0, col: 6, gen: true, num: 7},
// 	Cell{row: 0, col: 7, gen: true, num: 4},
// 	Cell{row: 0, col: 8, gen: true, num: 0},
// 	Cell{row: 1, col: 0, gen: true, num: 0},
// 	Cell{row: 1, col: 1, gen: true, num: 0},
// 	Cell{row: 1, col: 2, gen: true, num: 3},
// 	Cell{row: 1, col: 3, gen: true, num: 0},
// 	Cell{row: 1, col: 4, gen: true, num: 8},
// 	Cell{row: 1, col: 5, gen: true, num: 6},
// 	Cell{row: 1, col: 6, gen: true, num: 0},
// 	Cell{row: 1, col: 7, gen: true, num: 9},
// 	Cell{row: 1, col: 8, gen: true, num: 0},
// 	Cell{row: 2, col: 0, gen: true, num: 0},
// 	Cell{row: 2, col: 1, gen: true, num: 9},
// 	Cell{row: 2, col: 2, gen: true, num: 0},
// 	Cell{row: 2, col: 3, gen: true, num: 4},
// 	Cell{row: 2, col: 4, gen: true, num: 0},
// 	Cell{row: 2, col: 5, gen: true, num: 2},
// 	Cell{row: 2, col: 6, gen: true, num: 0},
// 	Cell{row: 2, col: 7, gen: true, num: 6},
// 	Cell{row: 2, col: 8, gen: true, num: 0},
// 	Cell{row: 3, col: 0, gen: true, num: 0},
// 	Cell{row: 3, col: 1, gen: true, num: 2},
// 	Cell{row: 3, col: 2, gen: true, num: 0},
// 	Cell{row: 3, col: 3, gen: true, num: 5},
// 	Cell{row: 3, col: 4, gen: true, num: 0},
// 	Cell{row: 3, col: 5, gen: true, num: 3},
// 	Cell{row: 3, col: 6, gen: true, num: 0},
// 	Cell{row: 3, col: 7, gen: true, num: 0},
// 	Cell{row: 3, col: 8, gen: true, num: 0},
// 	Cell{row: 4, col: 0, gen: true, num: 0},
// 	Cell{row: 4, col: 1, gen: true, num: 5},
// 	Cell{row: 4, col: 2, gen: true, num: 0},
// 	Cell{row: 4, col: 3, gen: true, num: 6},
// 	Cell{row: 4, col: 4, gen: true, num: 0},
// 	Cell{row: 4, col: 5, gen: true, num: 0},
// 	Cell{row: 4, col: 6, gen: true, num: 9},
// 	Cell{row: 4, col: 7, gen: true, num: 0},
// 	Cell{row: 4, col: 8, gen: true, num: 4},
// 	Cell{row: 5, col: 0, gen: true, num: 0},
// 	Cell{row: 5, col: 1, gen: true, num: 0},
// 	Cell{row: 5, col: 2, gen: true, num: 4},
// 	Cell{row: 5, col: 3, gen: true, num: 0},
// 	Cell{row: 5, col: 4, gen: true, num: 0},
// 	Cell{row: 5, col: 5, gen: true, num: 8},
// 	Cell{row: 5, col: 6, gen: true, num: 6},
// 	Cell{row: 5, col: 7, gen: true, num: 2},
// 	Cell{row: 5, col: 8, gen: true, num: 0},
// 	Cell{row: 6, col: 0, gen: true, num: 0},
// 	Cell{row: 6, col: 1, gen: true, num: 0},
// 	Cell{row: 6, col: 2, gen: true, num: 0},
// 	Cell{row: 6, col: 3, gen: true, num: 0},
// 	Cell{row: 6, col: 4, gen: true, num: 0},
// 	Cell{row: 6, col: 5, gen: true, num: 0},
// 	Cell{row: 6, col: 6, gen: true, num: 2},
// 	Cell{row: 6, col: 7, gen: true, num: 0},
// 	Cell{row: 6, col: 8, gen: true, num: 3},
// 	Cell{row: 7, col: 0, gen: true, num: 4},
// 	Cell{row: 7, col: 1, gen: true, num: 3},
// 	Cell{row: 7, col: 2, gen: true, num: 0},
// 	Cell{row: 7, col: 3, gen: true, num: 0},
// 	Cell{row: 7, col: 4, gen: true, num: 6},
// 	Cell{row: 7, col: 5, gen: true, num: 1},
// 	Cell{row: 7, col: 6, gen: true, num: 0},
// 	Cell{row: 7, col: 7, gen: true, num: 5},
// 	Cell{row: 7, col: 8, gen: true, num: 0},
// 	Cell{row: 8, col: 0, gen: true, num: 9},
// 	Cell{row: 8, col: 1, gen: true, num: 1},
// 	Cell{row: 8, col: 2, gen: true, num: 0},
// 	Cell{row: 8, col: 3, gen: true, num: 8},
// 	Cell{row: 8, col: 4, gen: true, num: 0},
// 	Cell{row: 8, col: 5, gen: true, num: 5},
// 	Cell{row: 8, col: 6, gen: true, num: 4},
// 	Cell{row: 8, col: 7, gen: true, num: 7},
// 	Cell{row: 8, col: 8, gen: true, num: 6},
// }

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
