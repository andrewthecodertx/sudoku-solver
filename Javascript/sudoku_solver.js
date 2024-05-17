var puzzle = [
  [7, 8, 0, 0, 2, 9, 0, 0, 3],
  [1, 4, 0, 0, 8, 0, 0, 6, 0],
  [0, 0, 0, 4, 0, 0, 1, 0, 0],
  [0, 0, 7, 0, 0, 0, 0, 0, 6],
  [5, 2, 0, 0, 0, 0, 0, 3, 1],
  [3, 0, 0, 0, 0, 0, 5, 0, 0],
  [0, 0, 8, 0, 0, 2, 0, 0, 0],
  [0, 7, 0, 0, 6, 0, 0, 1, 4],
  [6, 0, 0, 9, 4, 0, 0, 8, 2]
];

class Sudoku {
  constructor(puzzle) {
    this.sudoku = puzzle;
    this.solved = false;
  }

  isPossible(y, x, n) {
    const y0 = (Math.floor(y / 3) * 3);
    const x0 = (Math.floor(x / 3) * 3);

    for (let i = 0; i < 9; i++) {
      if (this.sudoku[y][i] == n) {
        return false;
      }
    }

    for (let i = 0; i < 9; i++) {
      if (this.sudoku[i][x] == n) {
        return false;
      }
    }

    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        if (this.sudoku[y0 + i][x0 + j] == n) {
          return false;
        }
      }
    }

    return true;
  };

  solve() {
    for (let y = 0; y < 9; y++) {
      for (let x = 0; x < 9; x++) {
        if (this.sudoku[y][x] == 0) {
          for (let n = 1; n <= 9; n++) {
            if (this.isPossible(y, x, n)) {
              this.sudoku[y][x] = n;
                if (this.solve()) {
                  return true;
                }

            	this.sudoku[y][x] = 0;
            }
        	}

        	return false;
        }
      }
    }

    return true;
  };
}

var s = new Sudoku(puzzle);
s.solve();

console.table(puzzle);
