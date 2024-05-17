let puzzle = 
  [|[|0; 0; 9; 0; 3; 0; 1; 0; 0|];
    [|4; 0; 0; 1; 0; 0; 0; 9; 0|];
    [|0; 5; 0; 0; 0; 4; 6; 0; 0|];
    [|0; 0; 0; 7; 0; 3; 5; 0; 0|];
    [|0; 8; 0; 0; 0; 0; 0; 2; 0|];
    [|0; 0; 1; 4; 0; 5; 0; 0; 0|];
    [|0; 0; 7; 2; 0; 0; 0; 5; 0|];
    [|0; 2; 0; 0; 0; 8; 0; 0; 7|];
    [|0; 0; 5; 0; 1; 0; 2; 0; 0|]|]

let is_possible x y n puzzle =
  let x0 = (x / 3) * 3 in
  let y0 = (y / 3) * 3 in

  let row_check = Array.exists ((=) n) puzzle.(y) in
  let col_check = Array.exists (fun row -> row.(x) = n) puzzle in
  let grid_check = 
    Array.exists (fun i -> 
      Array.exists (fun j -> 
        puzzle.(y0 + i).(x0 + j) = n) 
      [|0; 1; 2|]) [|0; 1; 2|] in
  
  not (row_check || col_check || grid_check)

let solve puzzle =
  let rec loop x y =
    if y = 9 then true
    else if puzzle.(y).(x) <> 0 then
      loop (if x = 8 then 0 else x + 1) (if x = 8 then y + 1 else y)
    else
      let rec try_n n =
        if n = 10 then false
        else if is_possible x y n puzzle then (
          puzzle.(y).(x) <- n;
          let (new_x, new_y) = if x = 8 then (0, y + 1) else (x + 1, y) in
          if loop new_x new_y then true
          else (puzzle.(y).(x) <- 0; try_n (n + 1))
        )
        else try_n (n + 1)
      in
      try_n 1
  in
  loop 0 0

let print_matrix puzzle =
  Array.iter (fun row ->
    Array.iter (fun node -> print_int node; print_string " ") row;
    print_newline ()) puzzle

let () =
  if solve puzzle then
    print_matrix puzzle
  else
    print_endline "no solution"
