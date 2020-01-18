package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// map of number location in cells
var positions map[int][]Position

// map of candidate cells for a number
var candidates map[int][]Position

// the puzzle
var cells [9][9]*Cell

func loadPuzzle() error {

	j, err := ioutil.ReadFile("puzzle.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, &cells)
	if err != nil {
		return err
	}

	if err := mapNumberPositions(); err != nil {
		return err
	}

	return nil
}

func mapNumberPositions() error {

	positions = make(map[int][]Position)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// iterate over all cells
			// map existing numbers to locations

			number := cells[i][j].Number

			if number < 0 || number > 9 {
				cells[i][j].selected = true
				return fmt.Errorf("invalid puzzle:%d is not a valid number", number)
			}

			cells[i][j].row = i
			cells[i][j].col = j
			if err := cells[i][j].setNumber(number); err != nil {
				return err
			}

			// register this position for this number
			var pos Position
			pos.row = i
			pos.col = j
			positions[number] = append(positions[number], pos)

		}
	}

	// above loop produces a map of number positions:
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("number %d found in: %#v", i, positions[i])
	// }

	return nil
}

// SetNum for a cell
func (c Cell) setNumber(num int) error {

	// check before setting this cell's number

	// but exclude 0's
	if num == 0 {
		return nil
	}

	// iterate over cell's row
	for i := 0; i < 9; i++ {

		//fmt.Printf("checking number: %d at row %d and column %d\n", num, c.row, i)

		if i == c.col { // exclude this cell
			//fmt.Printf("excluding col %d\n", i)
			continue
		}

		if num == cells[c.row][i].Number { // check if this number already exists in row
			cells[c.row][i].invalid = true
			return fmt.Errorf("error setting number '%d' for cell in row:%d; this number already exists at column:%d", num, c.row, i)

		}

	}

	// iterate over cell's column
	for i := 0; i < 9; i++ {

		if i == c.row {
			continue
		}

		if cells[i][c.col].Number == num { // check if this number already exists in column
			cells[i][c.col].invalid = true
			return fmt.Errorf("error setting number '%d' for cell in column:%d; this number already exists at row:%d", num, c.col, i)
		}
	}

	// iterate over cell's 9cell box
	cellpos := c.box()

	for i := cellpos.row; i < cellpos.row+3; i++ {
		for j := cellpos.col; j < cellpos.col+3; j++ {

			if c.row == i && c.col == j {
				continue // exclude this cell
			}

			if cells[i][j].Number == num {
				cells[i][j].invalid = true
				return fmt.Errorf("error setting number '%d' for cell in box row:%d and column:%d; this number already exists at box row:%d and col:%d", num, c.row, c.col, i, j)
			}
		}
	}

	return nil
}

// box returns a cell's first 9 cell box {position} (upper left most cell)
func (c Cell) box() Position {

	var pos Position

	// calculate the first cell, division, then multiply by 3
	crd3 := c.row / 3
	ccd3 := c.col / 3

	pos.row = crd3 * 3
	pos.col = ccd3 * 3

	return pos
}

// selects rows and columns containing a given number
func selectNumber(num int) error {

	// first find the cells (positions) containing this number
	numpos := positions[num]

	// select(highlight) rows and columns of those cells
	for _, v := range numpos {
		if err := selectCells(v.row, v.col); err != nil {
			fmt.Printf("error selecting row/column: %s\n", err)
		}
	}

	return nil
}

// selects a box, row and a column of a given cell position
func selectCells(row int, col int) error {

	// cannot select out of range
	if row < 0 || row > 8 {
		return fmt.Errorf("row out of range: %d", row)
	}

	if col < 0 || col > 8 {
		return fmt.Errorf("column out of range: %d", col)
	}

	// activate this cell to show number's position
	cells[row][col].active = true

	// there is no point in selecting an empty cell
	if cells[row][col].Number == 0 {
		return fmt.Errorf("%s", "cannot select an empty cell")
	}

	// first select the cells within that row
	for k := range cells[row] {
		cells[row][k].selected = true
	}

	// then select those within that column
	for k := range cells {
		cells[k][col].selected = true
	}

	// finally select those within that box
	box := cells[row][col].box()
	for k := box.row; k < box.row+3; k++ {
		for l := box.col; l < box.col+3; l++ {
			cells[k][l].selected = true
		}
	}

	return nil
}

// find/mark candidate cell positions for a number
func candidPos(num int) error {

	candidates = make(map[int][]Position)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			cells[i][j].marks = make([]int, 9)

			// every non-selected AND empty cell can accept a mark of this number
			if cells[i][j].selected == false {

				if cells[i][j].Number == 0 {

					// add mark to this cell as a candidate number
					cells[i][j].marks = append(cells[i][j].marks, num)

					// add this cell as a candidate position for this number
					var pos Position
					pos.row = i
					pos.col = j
					candidates[num] = append(candidates[num], pos)
				}

			} else {
				// remove from selected
				cells[i][j].selected = false
			}
		}
	}

	return nil
}

// Content prints a cell's number as a string
func (c Cell) Content() string {

	var number, color string

	color = "37" // default is white foreground color

	if c.invalid {
		color = "31"
	}

	if c.active {
		color = "35"
	}

	if c.solved {
		color = "1"
	}

	// zero-numbered cells shown as empty
	if c.Number == 0 {
		number = " "
	} else {
		number = fmt.Sprintf("%d", c.Number)
	}

	if c.candid {
		color = "42" // see structs for colors
	}

	return "\033[0;" + color + "m" + number + "\033[0m"
}

func clearConsole() {
	fmt.Println("\033[2J")
}

func clearSelected() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cells[i][j].selected = false
		}
	}
}

func clearActive() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cells[i][j].active = false
			cells[i][j].candid = false
		}
	}
}

func main() {

	// load puzzle from puzzle.json file
	if err := loadPuzzle(); err != nil {
		panic(err)
	}

	// show initial puzzle state
	fmt.Printf("new puzzle, difficulty: %s\n", difficulty())
	printBoard()

	var num int // search using crosshatch method
	num = getNumber()

	for {

		// select cells associated with this number;
		// rows, columns and 9cell boxes
		if err := selectNumber(num); err != nil {
			fmt.Println(err)
		}

		// from selections above, find candidate cells for this number
		if err := candidPos(num); err != nil {
			fmt.Println(err)
		}

		for _, v := range candidates[num] {
			//fmt.Printf("row %d, col %d\n", v.row, v.col)
			cells[v.row][v.col].candid = true
		}

		clearConsole()

		fmt.Printf("Number %d found in %d cells\n", num, len(positions[num]))
		fmt.Printf("%d candidate cells for number %d:\n", len(candidates[num]), num)

		printBoard()

		crosshatch(num)

		clearActive()
		// get a new number to crosshatch
		num = getNumber()
	}

}

// cross-hatching method:
// https://www.stolaf.edu/people/hansonr/sudoku/explain.htm#scanning
// algorithm explanation:
// starting with every number, search for all occurences
// save all positions of that number
// for each position, highlight all row and columns
// find all row and columns which that number is missing
func crosshatch(num int) {

	// easiest case; only one candidate
	if len(candidates[num]) == 1 {
		pos := candidates[num][0]
		putSolution(num, pos)

		return
	}

	// iterate over candid positions for num
	for i := 0; i < len(candidates[num]); i++ {

		// current position
		pos := candidates[num][i]

		// testing algorithm: surrounding cells
		// if no other candidate appears in the surrounding cells
		// assume this is the solution
		for _, v := range getSurroundingCells(pos) {
			if v.row != pos.row {
				if v.col != pos.col {
					// this is the solution?
					putSolution(num, pos)
				}
			}
		}

	}

}

// TODO
// get all cells next to a given position
func getSurroundingCells(pos Position) []Position {

	var positions []Position

	return positions
}

// puts a number to an empty cell and sets 'solved' to true
// 'solved' cell's numbers appear brighter
func putSolution(num int, pos Position) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("cell in row %d and col %d accepts the number %d, hit enter to complete\n", pos.row, pos.col, num)
	scanner.Scan()

	// the cell
	c := cells[pos.row][pos.col]

	// put solution and validate it
	if err := c.setNumber(num); err != nil {
		fmt.Println("error putting solution: ", err)
		return
	}

	cells[pos.row][pos.col].Number = num
	cells[pos.row][pos.col].solved = true
}

func getNumber() int {

	var num int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a number:")
		scanner.Scan()
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("\n")
			fmt.Printf("Must enter a number from 1 to 9\n")
			continue
		}

		if i > 9 || i < 1 {
			fmt.Printf("\n")
			fmt.Print("Must enter a number from 1 to 9\n")
			continue
		}

		num = i

		break

	}

	return num
}

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
