package main

import "fmt"

func main() {
	puzzle := [9][9]int{
		{0, 0, 9, 0, 3, 0, 1, 0, 0},
		{4, 0, 0, 1, 0, 0, 0, 9, 0},
		{0, 5, 0, 0, 0, 4, 6, 0, 0},
		{0, 0, 0, 7, 0, 3, 5, 0, 0},
		{0, 8, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 1, 4, 0, 5, 0, 0, 0},
		{0, 0, 7, 2, 0, 0, 0, 5, 0},
		{0, 2, 0, 0, 0, 8, 0, 0, 7},
		{0, 0, 5, 0, 1, 0, 2, 0, 0},
	}

	solve(&puzzle)
	printMatrix(&puzzle)
}

func isPossible(x, y, n int, puzzle *[9][9]int) bool {
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3

	for i := 0; i < 9; i++ {
		if puzzle[y][i] == n {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if puzzle[i][x] == n {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if puzzle[y0+i][x0+j] == n {
				return false
			}
		}
	}

	return true
}

func solve(puzzle *[9][9]int) bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if puzzle[y][x] == 0 {
				for n := 1; n <= 9; n++ {
					if isPossible(x, y, n, puzzle) {
						puzzle[y][x] = n
						if solve(puzzle) {
							return true
						}
						puzzle[y][x] = 0
					}
				}

				return false
			}
		}
	}

	return true
}

func printMatrix(puzzle *[9][9]int) {
	for _, row := range puzzle {
		for _, node := range row {
			fmt.Printf("%d ", node)
		}
		fmt.Println()
	}
}

