#!/bin/bash

# Set image name and container name
IMAGE_NAME="ascii-art-web-dockerize"
CONTAINER_NAME="ascii-art-web-dockerize-container"

# Build the Docker image
docker build -t $IMAGE_NAME .

# Run the Docker container
docker run -p 9000:9000 --name $CONTAINER_NAME $IMAGE_NAME
