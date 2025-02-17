# Use an official Golang runtime as a parent image
FROM golang:latest

RUN mkdir /app

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required third-party dependencies into the container
RUN go mod download

# Build the Go application
RUN go build -o order-service ./cmd/main/main.go

# Expose port 3051 to the outside world
EXPOSE 3052

# Command to run the executable
CMD ["./order-service"]
