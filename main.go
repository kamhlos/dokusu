package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	num int
	row int
	col int
}

type pos struct {
	row int
	col int
}

type Cell struct {
	num      int
	row      int
	col      int
	gen      bool  // auto-generated numbers have a default color
	valid    bool  // those have a red or green color
	selected bool  // selected cells have a magenta color
	marks    []int // other num candidates (if num = 0)
}

// map of number location in cells
var positions map[int][]pos

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

func mapNumberPositions() error {

	positions = make(map[int][]pos)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// iterate over all cells
			// map existing numbers to locations
			var position pos
			position.row = i
			position.col = j

			positions[cells[i][j].num] = append(positions[cells[i][j].num], position)

		}
	}

	// above loop produces a map of number positions:

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

	return nil
}

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

// selects rows and columns containing a given number
func selectNumber(num int) error {

	// first find the cells (positions) containing this number
	numpos := positions[num]

	fmt.Printf("number %d found in %d cells\n", num, len(numpos))

	// select(highlight) rows and columns of those cells
	for _, v := range numpos {
		if err := selectCells(v.row, v.col); err != nil {
			fmt.Printf("error selecting row/column: %s\n", err)
		}
	}

	return nil
}

// selects a row and a column of given cell position
func selectCells(row int, col int) error {

	// cannot select out of range
	if row < 0 || row > 8 {
		return fmt.Errorf("row out of range: %d", row)
	}

	if col < 0 || col > 8 {
		return fmt.Errorf("column out of range: %d", col)
	}

	// first select this cell
	cells[row][col].selected = true

	// there is no point in selecting an empty cell
	if cells[row][col].num == 0 {
		return fmt.Errorf("%s", "cannot select an empty cell")
	}

	// first select the cells within that row
	for k, _ := range cells[row] {
		cells[row][k].selected = true
	}

	// then select those within that column
	for k, _ := range cells {
		cells[k][col].selected = true
	}

	return nil
}

// Mark a possible candidate number for a cell
func markCell(row int, col int, num int) {

}

func (c Cell) Content() string {

	var number string
	color := "97" // default is white foreground color

	if c.selected {
		color = "96" // cyan foreground color
	}

	// zero-numbered cells shown as empty
	if c.num == 0 {
		number = " "
	} else {
		number = fmt.Sprintf("%d", c.num)
	}

	if !c.gen {
		if c.valid {
			color = "92"
		} else {
			color = "31"
		}
	}

	return "\033[0;" + color + "m" + number + "\033[0m"
}

func main() {

	crosshatch()
	// randomCells()

	// printBoard()

	// printBoard()

	// fmt.Printf("\n\n")
}

// difficulty measured by the count of 0's;
// > 35 considered easy, < 25 hard
func difficulty() string {

	if len(positions[0]) > 35 {
		return "easy"
	}

	if len(positions[0]) < 25 {
		return "hard"
	}

	return fmt.Sprintf("count of 0's: %d", len(positions[0]))
}

// cross-hatching method:
// https://www.stolaf.edu/people/hansonr/sudoku/explain.htm#scanning
// algorithm explanation:
// starting with every digit, search for all occurences
// save all positions of that digit
// for each position, highlight all row and columns
// find all row and columns which that digit is missing
func crosshatch() {

	var num int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a number for cross hatch:")
		scanner.Scan()
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("\n")
			fmt.Printf("1 to 9 only\n")
			continue
		}

		if i > 9 || i < 1 {
			fmt.Printf("\n")
			fmt.Print("1 to 9 only\n")
			continue
		}

		num = i

		break

	}

	// Use collected inputs
	fmt.Println(num)

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

	fmt.Println("\033[2J")

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
	printNumberRow(8)

	// lower border
	fmt.Printf("\u2517\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u253B\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u2537\u2501\u2501\u2501\u251B\n")

	// END third row of boxes

	fmt.Printf("\n\n")
}

func printNumberRow(n int) {

	// replace 2502 with 250A or 2506 for vertical lines as cells separators

	fmt.Printf("\u2503 %s \u2502 %s \u2502 %s \u2503", cells[n][0].Content(), cells[n][1].Content(), cells[n][2].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503", cells[n][3].Content(), cells[n][4].Content(), cells[n][5].Content())
	fmt.Printf(" %s \u2502 %s \u2502 %s \u2503\n", cells[n][6].Content(), cells[n][7].Content(), cells[n][8].Content())
}
