FROM golang:latest AS builder
WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

FROM alpine:latest
WORKDIR /app/

COPY --from=builder /app/app .

CMD ["./app"]
