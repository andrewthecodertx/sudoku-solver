class Sudoku:
    def __init__(self, puzzle):
        if self.is_valid_puzzle(puzzle):
            self.sudoku = puzzle
            self.solve()
        else:
            raise ValueError('invalid sudoku puzzle')

    @staticmethod
    def is_valid_puzzle(puzzle):
        if len(puzzle) != 9:
            return False
        for row in puzzle:
            if len(row) != 9:
                return False
            for cell in row:
                if not isinstance(cell, int) or not 0 <= cell <= 9:
                    return False
        return True

    def is_possible(self, x, y, n):
        for i in range(9):
            if self.sudoku[y][i] == n or self.sudoku[i][x] == n:
                return False

        x0 = (x // 3) * 3
        y0 = (y // 3) * 3
        for i in range(3):
            for j in range(3):
                if self.sudoku[y0 + i][x0 + j] == n:
                    return False
        return True

    def solve(self):
        for y in range(9):
            for x in range(9):
                if self.sudoku[y][x] == 0:
                    for n in range(1, 10):
                        if self.is_possible(x, y, n):
                            self.sudoku[y][x] = n
                            if self.solve():
                                return True
                            self.sudoku[y][x] = 0
                    return False
        return True

    def array_to_matrix(self):
        return '\n'.join([' '.join(map(str, row)) for row in self.sudoku])


# Example puzzle
puzzle = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 2, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 2, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0]
]

try:
    s = Sudoku(puzzle)
    print(s.array_to_matrix())
except ValueError as e:
    print(e)
