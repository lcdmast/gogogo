package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, colums = 9, 9

type Grid [rows][colums]int8

func (g *Grid) Set(row, colum int, digit int8) error {
	if !inBounds(row, colum) {
		return errors.New("out of bounds")
	}
	g[row][colum] = digit
	return nil
}

func inBounds(row, colum int) bool {
	if row < 0 || row >= rows {
		return false
	}

	if colum < 0 || colum >= colums {
		return false
	}

	return true
}

func main() {
	var g Grid
	err := g.Set(8, 0, 5)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		os.Exit(1)
	}
}
