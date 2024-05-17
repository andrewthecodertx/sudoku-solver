FROM gcc:latest

WORKDIR /usr/src/app

COPY . .

RUN g++ -o sudoku_solver sudoku_solver.cpp

CMD ["./sudoku_solver"]
