# Build Stage
FROM golang:1.21.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Golang application
RUN go build -o main main.go

# Final Stage
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
