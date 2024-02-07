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

# Copy the rest of the source code
COPY . .

# Adjust the WORKDIR to the specific directory where the main.go is located if it's not in the root
WORKDIR /app/producer

# Ensure all dependencies are downloaded
RUN go mod tidy

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/producer .

# Final stage: Create the runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the compiled binary from the second stage
COPY --from=gobuild /app/build/producer ./producer

EXPOSE 8080
CMD ["./producer"]
