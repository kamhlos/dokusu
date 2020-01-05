package main

import (
	"fmt"
	"testing"
)

func TestPrintBoard(t *testing.T) {

	err := mapNumberPositions()

	if err != nil {
		fmt.Println("puzzle is invalid")
	} else {
		fmt.Println("Difficulty: ", difficulty())
		fmt.Println("0's count: ", len(positions[0]))
		printBoard()
	}

}

func TestRandomCells(t *testing.T) {

	// var locations []map[int]string

	// testing cross-hatch method

	// debug
	// fmt.Printf("------------------------ DEBUG -------------------------\n")
	// fmt.Printf("%#v\n", positions)
	// fmt.Printf("------------------------ DEBUG -------------------------\n")

}

func TestHighlightRow(t *testing.T) {

	// for _, v := range positions[0] {
	// 	fmt.Println("0 is found in ", v)
	// }
}

// only used to generate board code
func TestPrintCellVar(t *testing.T) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("[9]Cell{row: %d, col: %d, gen: true, num: }\n", i, j)
		}

	}
}
