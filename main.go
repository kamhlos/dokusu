package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Cell struct {
	num   int
	gen   bool // auto-generated numbers have a default color
	valid bool // those have a red or green color
}

var numpos [8]Cell // where else this number is found

var cells = [9][9]Cell{
	[9]Cell{
		Cell{num: 8, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 5, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 9, gen: true},
		Cell{num: 7, gen: true},
		Cell{num: 4, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 3, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 8, gen: true},
		Cell{num: 6, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 9, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 9, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 4, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 2, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 6, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 2, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 5, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 3, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 5, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 6, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 9, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 4, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 4, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 8, gen: true},
		Cell{num: 6, gen: true},
		Cell{num: 2, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 2, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 3, gen: true},
	},
	[9]Cell{
		Cell{num: 4, gen: true},
		Cell{num: 3, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 6, gen: true},
		Cell{num: 1, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 5, gen: true},
		Cell{num: 0, gen: true},
	},
	[9]Cell{
		Cell{num: 9, gen: true},
		Cell{num: 1, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 8, gen: true},
		Cell{num: 0, gen: true},
		Cell{num: 5, gen: true},
		Cell{num: 4, gen: true},
		Cell{num: 7, gen: true},
		Cell{num: 6, gen: true},
	},
}

func (c Cell) Content() string {
	if c.num == 0 {
		return " " // zero-numbered cells shown as empty
	}

	color := "97" // white color for auto-generated numbers

	if !c.gen {
		if c.valid {
			color = "92"
		} else {
			color = "31"
		}
	}

	return "\033[0;" + color + "m" + fmt.Sprintf("%d", c.num) + "\033[0m"
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

	printBoard()

	printBoard()

	fmt.Printf("\n\n")
}

// cross-hatching method:
// https://www.stolaf.edu/people/hansonr/sudoku/explain.htm#scanning
// algorithm explanation:
// starting with every digit, search for all occurences
// save all positions of that digit
// for each position, highlight all row and columns
// find all row and columns which that digit is missing
func crosshatch() {

	location := make(map[int][]string)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// iterate over all cells
			// map existing numbers to locations
			location[cells[i][j].num] = append(location[cells[i][j].num], fmt.Sprintf("row:%d", i)+fmt.Sprintf("col:%d", j))

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
			cells[i][j].num = r.Intn(9) + 1
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

func printNumberLine(cells [9]Cell) {

	// replace 2502 with 250A or 2506 for vertical lines

	fmt.Printf("\u2503 %s \u2502 %s \u2502 %s \u2503", cells[0].Content(), cells[1].Content(), cells[2].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503", cells[3].Content(), cells[4].Content(), cells[5].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503\n", cells[6].Content(), cells[7].Content(), cells[8].Content())
}
