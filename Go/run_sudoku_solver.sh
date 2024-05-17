#!/bin/bash

IMAGE_NAME="sudoku_solver"

echo "Building the Docker image..."
docker build -t $IMAGE_NAME .

if [ $? -eq 0 ]; then
	echo "Docker image built successfully."

	echo "Running the Docker container..."
	docker run --rm $IMAGE_NAME

	if [ $? -eq 0 ]; then
		echo "Docker container ran successfully."
	else
		echo "Error: Docker container failed to run."
	fi
else
	echo "Error: Docker image build failed."
fi
