# Use an official Go runtime as a parent image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

COPY . .

# Copy the Go application source code into the container
COPY ./go.mod ./
COPY ./go.sum ./

#Running go mod download before copying the source code will ensure that the dependencies are downloaded only when the go.mod or go.sum files change.
RUN go mod download

# Build the Go application inside the container
RUN go build -o myapp

# Command to run the Go application when the container starts
CMD ["go", "run", "main.go"]
