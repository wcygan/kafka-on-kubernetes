# First stage: Build the Protobuf files using Buf
FROM bufbuild/buf:latest as bufbuild
WORKDIR /app
COPY . .
RUN buf generate proto

# Second stage: Build the Go binary
FROM golang:1.21 as gobuild
WORKDIR /app

# Copy the generated files from the previous stage
COPY --from=bufbuild /app/generated /app/generated

# Adjust the WORKDIR to the specific directory where the main.go is located if it's not in the root
WORKDIR /app/consumer

# Copy the go.mod and go.sum files
COPY consumer/go.mod consumer/go.sum ./

# Download all dependencies. They will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Modify this COPY command to only copy the files from the host's consumer folder to the container's consumer folder
COPY consumer/ ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/consumer .

# Final stage: Create the runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the compiled binary from the second stage
COPY --from=gobuild /app/build/consumer ./consumer

EXPOSE 8080
CMD ["./consumer"]