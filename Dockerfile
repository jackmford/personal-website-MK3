# Use a multi-stage build to keep the final image small
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/web/main.go

# Final lightweight image
FROM alpine:latest

# Set working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 4000

# Run the application
CMD ["./app"]

