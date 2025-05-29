# Use official Go image to build the application
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy



# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o great-thanos-app /app/cmd/main.go

# Use a minimal base image to run the application
FROM alpine:latest

WORKDIR /

# Copy the binary from the builder image
COPY --from=builder /app/great-thanos-app .

# Copy .env file if it exists
COPY .env .env

# Expose the port the app will run on
EXPOSE 8080

# Run the application
CMD ["./great-thanos-app"]
