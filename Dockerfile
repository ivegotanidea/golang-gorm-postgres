# Start by pulling the Go image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go mod tidy

RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

RUN apk update && apk add curl

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/main /app/main
COPY app.env .
# Expose port 8080 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["/app/main"]
