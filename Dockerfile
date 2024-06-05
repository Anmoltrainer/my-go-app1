# Use the official Golang image as the build stage
FROM golang:1.20 as builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy go.mod and go.sum files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o my-app .

# Use a minimal base image for the final container
FROM alpine:3.18

# Install required packages
RUN apk add --no-cache ca-certificates

# Set the working directory inside the final container
WORKDIR /app

# Copy the built Go application from the builder stage
COPY --from=builder /app/my-app .

# Ensure the application is executable
RUN chmod +x my-app

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./my-app"]
