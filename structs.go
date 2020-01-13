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
	initial  bool  // initiallly given numbers have a default color
	valid    bool  // those have a red or green color
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
	Reset      = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)
