# First stage: Build the Protobuf files using Buf
FROM bufbuild/buf:latest as bufbuild
WORKDIR /app
COPY . .
RUN buf generate proto

# Second stage: Build the Go binary
FROM golang:1.21 as gobuild
WORKDIR /app/producer
# Copy the generated files from the previous stage
COPY --from=bufbuild /app .
# Copy the rest of the Go source code
COPY . ..
RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/producer .

# Final stage: Create the runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
# Copy the compiled binary from the second stage
COPY --from=gobuild /app/build/producer ./build/producer
EXPOSE 8080
CMD ["./build/producer"]