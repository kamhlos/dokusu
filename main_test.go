package main

import (
	"testing"
)

func TestMapBoxes(t *testing.T) {

	// if err := loadPuzzle(); err != nil {
	// 	panic(err)
	// }

	// mapBoxes()

	// for k, v := range boxes {
	// 	fmt.Printf("cells in box %d%d :", k.row, k.col)
	// 	fmt.Printf("%d, \n", v)
	// }
}

func TestMapCandidsToBoxes(t *testing.T) {

	// if err := loadPuzzle(); err != nil {
	// 	panic(err)
	// }

	// mapBoxes()

	// num := 1

	// if err := selectNumber(num); err != nil {
	// 	fmt.Println(err)
	// }

	// if err := candidPos(num); err != nil {
	// 	fmt.Println(err)
	// }

	// for _, v := range candidates[num] {
	// 	cells[v.row][v.col].candid = true
	// }

	// if _, err := crosshatch(num); err != nil {
	// 	fmt.Println(err)
	// }

	// boxCandids := make(map[Position][]Position)

	// // map each candidate position to a box
	// for _, pos := range candidates[num] {

	// 	// this cell's box (identified by it's starting position)
	// 	box := cells[pos.row][pos.col].box()

	// 	// put this candidate position in the map
	// 	boxCandids[box] = append(boxCandids[box], pos)
	// }

	// fmt.Printf("number %d candidates found in %d boxes:\n", num, len(boxCandids))
	// for k, v := range boxCandids {
	// 	fmt.Printf("%d%d box has %d candidates: ", k.row, k.col, len(v))
	// 	for _, pos := range v {
	// 		fmt.Printf("%d%d, ", pos.row, pos.col)
	// 	}
	// 	fmt.Printf("\n")
	// }

}

func TestMarks(t *testing.T) {

	// if err := loadPuzzle(); err != nil {
	// 	panic(err)
	// }

	// clearConsole()

	// fmt.Printf("new puzzle, difficulty: %s\n", difficulty())

	// for i := 1; i < 10; i++ {

	// 	if err := selectNumber(i); err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	if err := candidPos(i); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	// show all marks
	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {

	// 		fmt.Printf("marks for cell: [%d][%d]: ", i, j)

	// 		for _, v := range cells[i][j].marks {
	// 			fmt.Printf("%d, ", v)
	// 		}

	// 		fmt.Printf("\n")

	// 	}
	// }

	// show all candidates
	// for i := 1; i < 10; i++ {

	// 	fmt.Printf("%d candidates for number %d: ", len(candidates[i]), i)

	// 	for _, v := range candidates[i] {
	// 		fmt.Printf("[%d][%d], ", v.row, v.col)
	// 	}

	// 	fmt.Printf("\n")

	// }

}

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

func TestInvalidPuzzle(t *testing.T) {

	// if err := mapNumberPositions(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// printBoard()

	// // first validate the given puzzle
	// if err := validPuzzle(); err != nil {
	// 	fmt.Print(err)
	// 	return
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

func TestSelectNumber(t *testing.T) {

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

}

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
