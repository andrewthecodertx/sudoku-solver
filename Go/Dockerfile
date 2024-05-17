FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go build -o sudoku_solver sudoku_solver.go

CMD ["./sudoku_solver"]
