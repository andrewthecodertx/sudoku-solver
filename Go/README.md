# Sudoku Solver in Go

This repository contains a Sudoku solver program written in Go. The Dockerfile and shell script provided will help you build and run the program without needing to install any build tools other than Docker.

## Prerequisites

- Docker installed on your system.

## Files

- `sudoku_solver.go`: The Go source file for the Sudoku solver.
- `Dockerfile`: The Dockerfile to build the Docker image.
- `run_sudoku_solver.sh`: The shell script to build and run the Docker container.

## Usage

1. Clone the repository and navigate to the project directory.

    ```sh
    git clone <repository_url>
    cd <repository_directory>
    ```

2. Make the shell script executable.

    ```sh
    chmod +x run_sudoku_solver.sh
    ```

3. Run the shell script to build the Docker image and run the container.

    ```sh
    ./run_sudoku_solver.sh
    ```

This script will handle building the Docker image and running the container, executing the Sudoku solver program.

## Notes

- Ensure that `sudoku_solver.go`, `Dockerfile`, and `run_sudoku_solver.sh` are in the same directory.
- The script will print error messages if the build or run process fails.
