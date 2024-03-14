FROM golang:1.22.1-alpine3.19 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# RUN go mod tidy
RUN go build -ldflags="-s -w" -o api .

FROM scratch

COPY --from=builder ["/build/api", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/api"]