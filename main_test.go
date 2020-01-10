package main

import "testing"

func TestPrintBoard(t *testing.T) {

	// err := mapNumberPositions()

	// if err != nil {
	// 	fmt.Println("puzzle is invalid")
	// } else {
	// 	fmt.Println("Difficulty: ", difficulty())
	// 	fmt.Println("0's count: ", len(positions[0]))
	// 	// cells[0][1].num = 9
	// 	// cells[0][1].gen = false
	// 	// cells[0][1].valid = true
	// 	// cells[0][1].selected = true
	// 	//printBoard()
	// }

}

func TestRandomCells(t *testing.T) {

	// var locations []map[int]string

	// testing cross-hatch method

	// debug
	// fmt.Printf("------------------------ DEBUG -------------------------\n")
	// fmt.Printf("%#v\n", positions)
	// fmt.Printf("------------------------ DEBUG -------------------------\n")

}

// func TestSelectNumber(t *testing.T) {

// 	err := selectNumber(num)
// 	if err != nil {
// 		fmt.Printf("error selecting rows/cols for number: %d\n", num)
// 	}

// 	// find empty non-selected cells
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if cells[i][j].selected == false {
// 				if cells[i][j].num == 0 {
// 					fmt.Printf("possible candidate for %d: row %d, col %d\n", num, i, j)
// 					markCell(i, j, num)
// 				}
// 			}
// 		}

// 	}

// 	printBoard()
// }

func TestSelectCells(t *testing.T) {

	// for a given cell cells[n][n] select a row and a column
	// if err := selectCells(4, 8); err != nil {
	// 	fmt.Printf("error selecting row/column: %s\n", err)
	// }

	// printBoard()

}

// only used to generate board code
// func TestPrintCellVar(t *testing.T) {
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			fmt.Printf("[9]Cell{row: %d, col: %d, gen: true, num: }\n", i, j)
// 		}

// 	}
// }
