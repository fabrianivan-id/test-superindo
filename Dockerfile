# Use an official Golang image as a base image for building the app
FROM golang:1.23.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . ./

# Build the application binary
RUN go build -o main .

# Use a lightweight image for the final stage
FROM alpine:latest

# Install required tools (e.g., for MySQL or Redis connectivity)
RUN apk --no-cache add ca-certificates tzdata

# Set the working directory in the final container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy migration files if needed
COPY /migrations /migrations

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]
