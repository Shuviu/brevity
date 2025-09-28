# Stage 1: The Builder
FROM golang:1.25-alpine AS builder

# Add go-sqlite dependency
RUN apk add --no-cache gcc libc-dev

WORKDIR /src

# Copy go.mod and go.sum from local src directory to the containers workdir directory
COPY ./src/go.mod ./
COPY ./src/go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY ./src/ .

RUN go build -o /go-server ./cmd/httpserver/

# Stage 2: The Final Image
FROM alpine:latest

# The final image only needs the compiled binary
WORKDIR /app
COPY --from=builder /go-server .

EXPOSE 8080
CMD ["/app/go-server", ""]