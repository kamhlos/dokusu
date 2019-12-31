package main

import (
	"fmt"
	"testing"
)

// func TestPrintBoard(t *testing.T) {
// 	printBoard(cells)
// }

func TestRandomCells(t *testing.T) {

	location := make(map[int][]string)

	// var locations []map[int]string

	// testing cross-hatch method
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// iterate over all cells
			// map existing numbers to locations
			location[cells[i][j].num] = append(location[cells[i][j].num], fmt.Sprintf("row:%d", i)+fmt.Sprintf("col:%d", j))

		}
	}

	// debug
	fmt.Printf("------------------------ DEBUG -------------------------\n")
	fmt.Printf("%#v\n", location[8])
	fmt.Printf("------------------------ DEBUG -------------------------\n")

}

func TestHighlightRow(t *testing.T) {

}
