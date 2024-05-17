package main

import (
  "fmt"
)

func main() {
  puzzle := [][]int{
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

  err := solve(puzzle)
  if err != nil {
    fmt.Println("failed to solve sudoku:", err)
    return
  }
  printMatrix(puzzle)
}

func isPossible(x, y, n int, puzzle [][]int) bool {
  x0 := (x / 3) * 3
  y0 := (y / 3) * 3

  for i := 0; i < 9; i++ {
    if puzzle[y][i] == n || puzzle[i][x] == n || puzzle[y0+i%3][x0+i/3] == n {
      return false
    }
  }
  return true
}

func solve(puzzle [][]int) error {
  for y := 0; y < 9; y++ {
    for x := 0; x < 9; x++ {
      if puzzle[y][x] == 0 {
        for n := 1; n <= 9; n++ {
          if isPossible(x, y, n, puzzle) {
            puzzle[y][x] = n
            if err := solve(puzzle); err == nil {
              return nil
            }
            puzzle[y][x] = 0
          }
        }
        return fmt.Errorf("no solution")
      }
    }
  }
  return nil
}

func printMatrix(puzzle [][]int) {
  for _, row := range puzzle {
    for _, node := range row {
      fmt.Printf("%d ", node)
    }
    fmt.Println()
  }
}
