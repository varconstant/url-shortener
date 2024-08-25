# Step 1: Build the Go app using the Makefile
FROM golang:1.23.0-alpine3.19 AS builder

# Install make and other necessary build tools
RUN apk add --no-cache make gcc libc-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application using the Makefile
RUN make build

# Step 2: Create a smaller image for running the application
FROM alpine:3.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/app .

# Expose port 9001 (or whatever port your app uses)
EXPOSE 9001

# Command to run the executable
CMD ["./app"]
