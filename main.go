package main

import (
	"fmt"
)

var cells = [9][9]string{
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
	[9]string{
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
		" ",
	},
}

var reset = "\033[0m"

func main() {

	cells[0][7] = "\033[1;31m5\033[0m"
	cells[0][7] = "\033[2;94m2\033[0m"
	cells[2][2] = printNumber(true, false, 1)
	cells[8][4] = printNumber(false, true, 9)
	cells[4][0] = printNumber(false, validateCell(cells[4][0]), 6)

	printBoard(cells)

	fmt.Printf("\n\n")

}

func validateCell(cell [9][9]string) bool {
	// TODO:
	// logic to validate each cell

	return false
}

func printNumber(gen bool, valid bool, num int) string {

	color := "97" // white color for auto-generated numbers

	if !gen {
		if valid {
			color = "92"
		} else {
			color = "31"
		}
	}

	return "\033[0;" + color + "m" + fmt.Sprintf("%d", num) + reset
}

func printBoard(cells [9][9]string) {

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
