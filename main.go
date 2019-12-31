package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	number    int
	generated bool // those have a default color
	valid     bool // those have a red or green color
}

var cells = [9][9]Cell{
	[9]Cell{
		"8",
		" ",
		"5",
		" ",
		" ",
		"9",
		"7",
		"4",
		" ",
	},
	[9]Cell{
		" ",
		" ",
		"3",
		" ",
		"8",
		"6",
		" ",
		"9",
		" ",
	},
	[9]Cell{
		" ",
		"9",
		" ",
		"4",
		" ",
		"2",
		" ",
		"6",
		" ",
	},
	[9]Cell{
		" ",
		"2",
		" ",
		"5",
		" ",
		"3",
		" ",
		" ",
		" ",
	},
	[9]Cell{
		" ",
		"5",
		" ",
		"6",
		" ",
		" ",
		"9",
		" ",
		"4",
	},
	[9]Cell{
		" ",
		" ",
		"4",
		" ",
		" ",
		"8",
		"6",
		"2",
		" ",
	},
	[9]Cell{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		"2",
		" ",
		"3",
	},
	[9]Cell{
		"4",
		"3",
		" ",
		" ",
		"6",
		"1",
		" ",
		"5",
		" ",
	},
	[9]Cell{
		"9",
		"1",
		" ",
		"8",
		" ",
		"5",
		"4",
		"7",
		"6",
	},
}

func (c Cell) Content() {
	if c.number == 0 {

	}
}

func main() {

	// cells[0][7] = "\033[1;31m5\033[0m"
	// cells[0][7] = "\033[2;94m2\033[0m"
	// cells[2][2] = printNumber(true, false, 1)
	// cells[8][4] = printNumber(false, true, 9)
	// cells[4][0] = printNumber(false, false, 6)

	// populate cells array
	// i.e generate a random puzzle to solve
	// randomCells()

	printBoard(cells)
	solve()

	printBoard(cells)

	fmt.Printf("\n\n")

}

func solve() {

	var crn int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			if cells[i][j] != " " {
				crn = cells[i][j].number
				highlight(i, j, crn)
				break
			}
		}
	}
}

// cross-hatching method:
// https://www.stolaf.edu/people/hansonr/sudoku/explain.htm#scanning
// algorithm explanation:
// starting with every digit, search for all occurences
// save all positions of that digit
// for each position, highlight all row and columns
// find all row and columns which that digit is missing
func crosshatch() {

	// locations := make(map[int][9]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// start with the number 1
			// locations[cells[i][j]] = append(, )
			if cells[i][j] == "1" {
				return
			}
		}
	}
}

func highlight(line int, column int, num int) {
	// cells[line][column] = printCell(false, true, num)
}

func randomCells() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cells[i][j] = fmt.Sprintf("%d", r.Intn(9)+1)
		}
	}

	// for i := 0; i < 9; i++ {
	// 	ec := r.Intn(10)
	// 	cells[i][ec] = " "
	// 	cells[i][ec+r.Intn(n)] = " "
	// }
}

func validateCell(cell [9][9]string) bool {
	// TODO:
	// logic to validate each cell

	return false
}

func printCell(gen bool, valid bool, num string) string {

	color := "97" // white color for auto-generated numbers

	if !gen {
		if valid {
			color = "92"
		} else {
			color = "31"
		}
	}

	return "\033[0;" + color + "m" + num + "\033[0m"
}

// prints the board with the cells contents if any
func printBoard() {

	// START first line of boxes

	// upper border
	fmt.Printf("\u250F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2533\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2533\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2513\n")

	// first line of numbers
	printNumberLine(cells[0])

	// first middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second line of numbers
	printNumberLine(cells[1])

	// second middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third line of numbers
	printNumberLine(cells[2])

	// lower border
	fmt.Printf("\u2523\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u252B\n")

	// END first line of boxes

	// REPEAT
	// START second line of boxes (no border)

	// first line of numbers
	printNumberLine(cells[3])

	// first middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second line of numbers
	printNumberLine(cells[4])

	// second middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third line of numbers
	printNumberLine(cells[5])

	// lower border
	fmt.Printf("\u2523\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u252B\n")

	// END second line of boxes

	// REPEAT
	// START third line of boxes (no border)

	// first line of numbers
	printNumberLine(cells[6])

	// first middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second line of numbers
	printNumberLine(cells[7])

	// second middle line
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third line of numbers
	printNumberLine(cells[8])

	// lower border
	fmt.Printf("\u2517\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u251B\n")

	// END third line of boxes

	fmt.Printf("\n\n")
}

func printNumberLine(cells [9]string) {

	// replace 2502 with 250A or 2506 for vertical lines

	fmt.Printf("\u2503 %s \u2502 %s \u2502 %s \u2503", cells[0], cells[1], cells[2])
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503", cells[3], cells[4], cells[5])
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503\n", cells[6], cells[7], cells[8])
}
