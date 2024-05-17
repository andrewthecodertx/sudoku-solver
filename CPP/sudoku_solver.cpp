#include <iostream>

bool isPossible(int x, int y, int n, int puzzle[9][9]) {
  int x0 = (x / 3) * 3;
  int y0 = (y / 3) * 3;

  for (int i = 0; i < 9; i++) {
    if (puzzle[y][i] == n || puzzle[i][x] == n ||
        puzzle[y0 + i / 3][x0 + i % 3] == n) {
      return false;
    }
  }

  return true;
}

bool solve(int puzzle[9][9]) {
  for (int y = 0; y < 9; y++) {
    for (int x = 0; x < 9; x++) {
      if (puzzle[y][x] == 0) {
        for (int n = 1; n <= 9; n++) {
          if (isPossible(x, y, n, puzzle)) {
            puzzle[y][x] = n;
            if (solve(puzzle)) {
              return true;
            }
            puzzle[y][x] = 0;
          }
        }
        return false;
      }
    }
  }
  return true;
}

void printMatrix(int puzzle[9][9]) {
  for (int y = 0; y < 9; y++) {
    for (int x = 0; x < 9; x++) {
      std::cout << puzzle[y][x] << " ";
    }
    std::cout << std::endl;
  }
}

int main() {
  int puzzle[9][9] = {{0, 0, 9, 0, 3, 0, 1, 0, 0}, {4, 0, 0, 1, 0, 0, 0, 9, 0},
                      {0, 5, 0, 0, 0, 4, 6, 0, 0}, {0, 0, 0, 7, 0, 3, 5, 0, 0},
                      {0, 8, 0, 0, 0, 0, 0, 2, 0}, {0, 0, 1, 4, 0, 5, 0, 0, 0},
                      {0, 0, 7, 2, 0, 0, 0, 5, 0}, {0, 2, 0, 0, 0, 8, 0, 0, 7},
                      {0, 0, 5, 0, 1, 0, 2, 0, 0}};

  solve(puzzle);
  printMatrix(puzzle);

  return 0;
}

