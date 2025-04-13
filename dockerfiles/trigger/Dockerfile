
FROM golang:1.21-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /argocd-sync .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /argocd-sync /argocd-sync

USER nonroot:nonroot

ENTRYPOINT ["/argocd-sync"]
