fn main() {
  let mut puzzle: [[i32; 9]; 9] = [
    [0, 0, 9, 0, 3, 0, 1, 0, 0],
    [4, 0, 0, 1, 0, 0, 0, 9, 0],
    [0, 5, 0, 0, 0, 4, 6, 0, 0],
    [0, 0, 0, 7, 0, 3, 5, 0, 0],
    [0, 8, 0, 0, 0, 0, 0, 2, 0],
    [0, 0, 1, 4, 0, 5, 0, 0, 0],
    [0, 0, 7, 2, 0, 0, 0, 5, 0],
    [0, 2, 0, 0, 0, 8, 0, 0, 7],
    [0, 0, 5, 0, 1, 0, 2, 0, 0],
  ];

  solve(&mut puzzle);
  array_to_matrix(&mut puzzle);
}

fn is_possible(x: usize, y: usize, n: i32, puzzle: &[[i32; 9]; 9]) -> bool {
  let x0 = (x / 3) * 3;
  let y0 = (y / 3) * 3;

  for i in 0..9 {
    if puzzle[y][i] == n || puzzle[i][x] == n || puzzle[y0 + i % 3][x0 + i / 3] == n {
      return false;
    }
  }
  true
}

fn solve(puzzle: &mut [[i32; 9]; 9]) -> bool {
  for y in 0..9 {
    for x in 0..9 {
      if puzzle[y][x] == 0 {
        for n in 1..=9 {
          if is_possible(x, y, n, puzzle) {
            puzzle[y][x] = n;
            if solve(puzzle) {
              return true;
            }
            puzzle[y][x] = 0;
          }
        }
        return false;
      }
    }
  }
  true
}

fn array_to_matrix(puzzle: &[[i32; 9]; 9]) {
  for row in puzzle {
    for node in row {
      print!("{} ", node);
    }
    print!("\n");
  }
}
