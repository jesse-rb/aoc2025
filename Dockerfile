FROM golang:1.25-alpine AS builder-go

ARG SERVICE_DIR
WORKDIR /app

COPY go.mod ./
COPY go.sum* ./
RUN go mod download

COPY . .

# Build the binary
RUN go build -o app ./${SERVICE_DIR}

# Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder-go /app/app .

CMD ["./app"]
