package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// map of number location in cells
var positions map[int][]pos

// the puzzle
var cells = [9][9]Cell{
	[9]Cell{
		Cell{num: 8, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 5, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 9, initial: true},
		Cell{num: 7, initial: true},
		Cell{num: 4, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 3, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 8, initial: true},
		Cell{num: 6, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 9, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 9, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 4, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 2, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 6, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 2, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 5, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 3, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 5, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 6, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 9, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 4, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 4, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 8, initial: true},
		Cell{num: 6, initial: true},
		Cell{num: 2, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 2, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 3, initial: true},
	},
	[9]Cell{
		Cell{num: 4, initial: true},
		Cell{num: 3, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 6, initial: true},
		Cell{num: 1, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 5, initial: true},
		Cell{num: 0, initial: true},
	},
	[9]Cell{
		Cell{num: 9, initial: true},
		Cell{num: 1, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 8, initial: true},
		Cell{num: 0, initial: true},
		Cell{num: 5, initial: true},
		Cell{num: 4, initial: true},
		Cell{num: 7, initial: true},
		Cell{num: 6, initial: true},
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

			number := cells[i][j].num

			if number < 0 || number > 9 {
				return fmt.Errorf("invalid puzzle:%d is not a valid number", number)
			}

			positions[number] = append(positions[number], position)

		}
	}

	// above loop produces a map of number positions:
	for i := 0; i < 10; i++ {
		fmt.Printf("number %d found in: %#v", i, positions[i])
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
	for k := range cells[row] {
		cells[row][k].selected = true
	}

	// then select those within that column
	for k := range cells {
		cells[k][col].selected = true
	}

	return nil
}

// mark a possible candidate number for a cell
func markCell(row int, col int, num int) {
	cells[row][col].marks = append(cells[row][col].marks, num)
}

// Content prints a cell's number as a string
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

	if !c.initial {
		if c.valid {
			color = "92"
		} else {
			color = "31"
		}
	}

	return "\033[0;" + color + "m" + number + "\033[0m"
}

func validPuzzle() error {

	// TODO:
	// validate current number positions
	// return error if puzzle is not valid
	// RULES for validation:
	// each row/column/9cell box has no more than one copy of each number of 1 to 9
	// for k, v := range positions {

	// 	// each number must appear only once in a row
	// 	// with the exception of 0
	// 	//
	// 	for pos := range v {
	// 		// debug
	// 		fmt.Printf("checking number: %d\n", pos)
	// 		fmt.Printf("%#v\n", v)
	// 		fmt.Printf("------------------------ DEBUG -------------------------\n")
	// 	}

	// }

	return nil
}
func main() {

	// first validate the given puzzle
	if err := validPuzzle(); err != nil {
		fmt.Print(err)
		return
	}

	//num := getNumber()
	// crosshatch()
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
// starting with every number, search for all occurences
// save all positions of that number
// for each position, highlight all row and columns
// find all row and columns which that number is missing
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
