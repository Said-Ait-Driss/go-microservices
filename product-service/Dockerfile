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
RUN go build -o product-service ./cmd/main/main.go

# Expose port 3030 to the outside world
EXPOSE 3040

# Command to run the executable
CMD ["./product-service"]
