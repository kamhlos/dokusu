package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Number struct {
	num int
	row int
	col int
}

type Cell struct {
	num      int
	row      int
	col      int
	gen      bool  // auto-generated numbers have a default color
	valid    bool  // those have a red or green color
	selected bool  // selected cells have a yellow color
	marks    []int // other num candidates (if num = 0)
}

// map of number location in cells
var positions map[int][]string

var board = [81]Cell{
	Cell{row: 0, col: 0, gen: true, num: 8},
	Cell{row: 0, col: 1, gen: true, num: 0},
	Cell{row: 0, col: 2, gen: true, num: 5},
	Cell{row: 0, col: 3, gen: true, num: 0},
	Cell{row: 0, col: 4, gen: true, num: 0},
	Cell{row: 0, col: 5, gen: true, num: 9},
	Cell{row: 0, col: 6, gen: true, num: 7},
	Cell{row: 0, col: 7, gen: true, num: 4},
	Cell{row: 0, col: 8, gen: true, num: 0},
	Cell{row: 1, col: 0, gen: true, num: 0},
	Cell{row: 1, col: 1, gen: true, num: 0},
	Cell{row: 1, col: 2, gen: true, num: 3},
	Cell{row: 1, col: 3, gen: true, num: 0},
	Cell{row: 1, col: 4, gen: true, num: 8},
	Cell{row: 1, col: 5, gen: true, num: 6},
	Cell{row: 1, col: 6, gen: true, num: 0},
	Cell{row: 1, col: 7, gen: true, num: 9},
	Cell{row: 1, col: 8, gen: true, num: 0},
	Cell{row: 2, col: 0, gen: true, num: 0},
	Cell{row: 2, col: 1, gen: true, num: 9},
	Cell{row: 2, col: 2, gen: true, num: 0},
	Cell{row: 2, col: 3, gen: true, num: 4},
	Cell{row: 2, col: 4, gen: true, num: 0},
	Cell{row: 2, col: 5, gen: true, num: 2},
	Cell{row: 2, col: 6, gen: true, num: 0},
	Cell{row: 2, col: 7, gen: true, num: 6},
	Cell{row: 2, col: 8, gen: true, num: 0},
	Cell{row: 3, col: 0, gen: true, num: 0},
	Cell{row: 3, col: 1, gen: true, num: 2},
	Cell{row: 3, col: 2, gen: true, num: 0},
	Cell{row: 3, col: 3, gen: true, num: 5},
	Cell{row: 3, col: 4, gen: true, num: 0},
	Cell{row: 3, col: 5, gen: true, num: 3},
	Cell{row: 3, col: 6, gen: true, num: 0},
	Cell{row: 3, col: 7, gen: true, num: 0},
	Cell{row: 3, col: 8, gen: true, num: 0},
	Cell{row: 4, col: 0, gen: true, num: 0},
	Cell{row: 4, col: 1, gen: true, num: 5},
	Cell{row: 4, col: 2, gen: true, num: 0},
	Cell{row: 4, col: 3, gen: true, num: 6},
	Cell{row: 4, col: 4, gen: true, num: 0},
	Cell{row: 4, col: 5, gen: true, num: 0},
	Cell{row: 4, col: 6, gen: true, num: 9},
	Cell{row: 4, col: 7, gen: true, num: 0},
	Cell{row: 4, col: 8, gen: true, num: 4},
	Cell{row: 5, col: 0, gen: true, num: 0},
	Cell{row: 5, col: 1, gen: true, num: 0},
	Cell{row: 5, col: 2, gen: true, num: 4},
	Cell{row: 5, col: 3, gen: true, num: 0},
	Cell{row: 5, col: 4, gen: true, num: 0},
	Cell{row: 5, col: 5, gen: true, num: 8},
	Cell{row: 5, col: 6, gen: true, num: 6},
	Cell{row: 5, col: 7, gen: true, num: 2},
	Cell{row: 5, col: 8, gen: true, num: 0},
	Cell{row: 6, col: 0, gen: true, num: 0},
	Cell{row: 6, col: 1, gen: true, num: 0},
	Cell{row: 6, col: 2, gen: true, num: 0},
	Cell{row: 6, col: 3, gen: true, num: 0},
	Cell{row: 6, col: 4, gen: true, num: 0},
	Cell{row: 6, col: 5, gen: true, num: 0},
	Cell{row: 6, col: 6, gen: true, num: 2},
	Cell{row: 6, col: 7, gen: true, num: 0},
	Cell{row: 6, col: 8, gen: true, num: 3},
	Cell{row: 7, col: 0, gen: true, num: 4},
	Cell{row: 7, col: 1, gen: true, num: 3},
	Cell{row: 7, col: 2, gen: true, num: 0},
	Cell{row: 7, col: 3, gen: true, num: 0},
	Cell{row: 7, col: 4, gen: true, num: 6},
	Cell{row: 7, col: 5, gen: true, num: 1},
	Cell{row: 7, col: 6, gen: true, num: 0},
	Cell{row: 7, col: 7, gen: true, num: 5},
	Cell{row: 7, col: 8, gen: true, num: 0},
	Cell{row: 8, col: 0, gen: true, num: 9},
	Cell{row: 8, col: 1, gen: true, num: 1},
	Cell{row: 8, col: 2, gen: true, num: 0},
	Cell{row: 8, col: 3, gen: true, num: 8},
	Cell{row: 8, col: 4, gen: true, num: 0},
	Cell{row: 8, col: 5, gen: true, num: 5},
	Cell{row: 8, col: 6, gen: true, num: 4},
	Cell{row: 8, col: 7, gen: true, num: 7},
	Cell{row: 8, col: 8, gen: true, num: 6},
}

// var cells = [9][9]Cell{
// 	[9]Cell{
// 		Cell{num: 8, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 5, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 9, gen: true},
// 		Cell{num: 7, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 3, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 8, gen: true},
// 		Cell{num: 6, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 9, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 9, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 2, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 6, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 2, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 5, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 3, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 5, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 6, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 9, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 8, gen: true},
// 		Cell{num: 6, gen: true},
// 		Cell{num: 2, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 2, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 3, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 3, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 6, gen: true},
// 		Cell{num: 1, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 5, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 9, gen: true},
// 		Cell{num: 1, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 8, gen: true},
// 		Cell{num: 0, gen: true},
// 		Cell{num: 5, gen: true},
// 		Cell{num: 4, gen: true},
// 		Cell{num: 7, gen: true},
// 		Cell{num: 6, gen: true},
// 	},
// }

// func mapNumberPositions() error {

// 	positions = make(map[int][]string)

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			// iterate over all cells
// 			// map existing numbers to locations

// 			positions[cells[i][j].num] = append(positions[cells[i][j].num], fmt.Sprintf("%d:", i)+fmt.Sprintf("%d", j))

// 		}
// 	}

// 	// above loop produces a map of number positions:
	
// 		map[int][]string{
// 			0:[]string{
// 				"0:1", "0:3", "0:4", "0:8", "1:0", "1:1", "1:3", "1:6", "1:8", "2:0", "2:2", "2:4", "2:6", "2:8", "3:0", "3:2", "3:4", "3:6", "3:7", "3:8", "4:0", "4:2", "4:4", "4:5", "4:7", "5:0", "5:1", "5:3", "5:4", "5:8", "6:0", "6:1", "6:2", "6:3", "6:4", "6:5", "6:7", "7:2", "7:3", "7:6", "7:8", "8:2", "8:4"
// 			},
// 			1:[]string{
// 				"7:5", "8:1"
// 			},
// 			2:[]string{
// 				"2:5", "3:1", "5:7", "6:6"
// 			},
// 			3:[]string{
// 				"1:2", "3:5", "6:8", "7:1"
// 			},
// 			4:[]string{
// 				"0:7", "2:3", "4:8", "5:2", "7:0", "8:6"
// 			},
// 			5:[]string{
// 				"0:2", "3:3", "4:1", "7:7", "8:5"
// 			},
// 			6:[]string{
// 				"1:5", "2:7", "4:3", "5:6", "7:4", "8:8"
// 			},
// 			7:[]string{
// 				"0:6", "8:7"
// 			},
// 			8:[]string{
// 				"0:0", "1:4", "5:5", "8:3"
// 			},
// 			9:[]string{"0:5", "1:7", "2:1", "4:6", "8:0"
// 			}
// 		}
	

// 	return nil
// }

func validatePositions() error {

	// TODO:
	// validate current number positions
	// return error if puzzle is not valid
	// RULES for validation:
	// each row/column/9cell box has no more than one copy of each number of 1 to 9
	for i := 1; i < 10; i++ {

	}

	return nil
}

func (c Cell) highlight() error {

	c.selected = true

	return nil
}

func (c Cell) Content() string {
	if c.num == 0 {
		return " " // zero-numbered cells shown as empty
	}

	color := "97" // white color for auto-generated numbers

	if !c.gen {
		if c.valid {
			if c.selected {
				color = "84"
			} else {
				color = "92"
			}
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

	// printBoard()

	// printBoard()

	// fmt.Printf("\n\n")
}

// difficulty measured by the count of 0's;
// > 30 considered easy, < 25 probably hard
func difficulty() string {

	// 	if len(positions[0]) > 30 {
	// 		return "easy"
	// 	}

	// 	if len(positions[0]) < 25 {
	// 		return "hard"
	// 	}

	// 	return fmt.Sprintf("count of 0's: %d", len(positions[0]))
}

// cross-hatching method:
// https://www.stolaf.edu/people/hansonr/sudoku/explain.htm#scanning
// algorithm explanation:
// starting with every digit, search for all occurences
// save all positions of that digit
// for each position, highlight all row and columns
// find all row and columns which that digit is missing
func crosshatch() {

	// location := make(map[int][]string)

	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		// iterate over all cells
	// 		// map existing numbers to locations
	// 		location[cells[i][j].num] = append(location[cells[i][j].num], fmt.Sprintf("row:%d", i)+fmt.Sprintf("col:%d", j))

	// 	}
	// }
}

func randomCells() {

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		cells[i][j].num = r.Intn(9) + 1
	// 	}
	// }

	// for i := 0; i < 9; i++ {
	// 	ec := r.Intn(10)
	// 	cells[i][ec] = " "
	// 	cells[i][ec+r.Intn(n)] = " "
	// }
}

// func validateCell(cell [9][9]string) bool {
// 	// TODO:
// 	// logic to validate each cell

// 	return false
// }

// prints the board with the cells contents if num not zero
func printBoard() {

	// START first row of boxes

	// upper border
	fmt.Printf("\u250F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2533\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2533\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u252F\u2501\u2501\u2501\u2513\n")

	// first row of numbers
	printNumberRow(0)

	// first middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second row of numbers
	printNumberRow(1)

	// second middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third row of numbers
	printNumberRow(2)

	// lower border
	fmt.Printf("\u2523\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u252B\n")

	// END first row of boxes

	// REPEAT
	// START second row of boxes (no border)

	// first row of numbers
	printNumberRow(3)

	// first middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second row of numbers
	printNumberRow(4)

	// second middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third row of numbers
	printNumberRow(5)

	// lower border
	fmt.Printf("\u2523\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u254B\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u253F\u2501\u2501\u2501\u252B\n")

	// END second row of boxes

	// REPEAT
	// START third row of boxes (no border)

	// first row of numbers
	printNumberRow(6)

	// first middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// second row of numbers
	printNumberRow(7)

	// second middle row
	fmt.Printf("\u2520\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2542\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2528\n")

	// third row of numbers
	printNumberRow(8])

	// lower border
	fmt.Printf("\u2517\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u251B\n")

	// END third row of boxes

	fmt.Printf("\n\n")
}

func printNumberRow(n int) {

	numtoprint = make(map[int]string)
	// replace 2502 with 250A or 2506 for vertical lines as cells separators
	for i := 0; i < 9; i++ {
		numtoprint[i] = Cell
	}

	fmt.Printf("\u2503 %s \u2502 %s \u2502 %s \u2503", Cell., cells[1].Content(), cells[2].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503", cells[3].Content(), cells[4].Content(), cells[5].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503\n", cells[6].Content(), cells[7].Content(), cells[8].Content())
}
