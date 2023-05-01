FROM golang:latest

# Set the working directory in the container
WORKDIR /AuthApp

# Copy the necessary files into the container
COPY . .
COPY go.mod go.sum ./
COPY main.go ./

# Download the Go modules
RUN go mod tidy

# Build the binary
RUN go build .

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./main"]