# Stage 1 - Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /webhook-receiver ./cmd/webhook/main.go

# Stage 2 - Final image
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /webhook-receiver /webhook-receiver

USER nonroot:nonroot

ENTRYPOINT ["/webhook-receiver"]
