FROM golang:1.23.5

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the binary from the cmd directory
RUN go build -o main ./cmd/main.go

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["/app/main"]
