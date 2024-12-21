# Fetch base image file
FROM golang:1.22.2-alpine

# Create a bash environment for the container
RUN apk add --no-cache bash

# Additional image meta data
LABEL version="1.0"
LABEL description="Image file for a website project that does ascii conversions"

# Specify the image working directory
WORKDIR /app

# Copy the rest base source code to the working directory
COPY . .

# Command to build an executable file
RUN go build -o ascii-art-web-dockerize . 

# Expose container to the outside world
EXPOSE 9000

# Run the executable file
CMD ["./ascii-art-web-dockerize"]