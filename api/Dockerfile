FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /api ./cmd/

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /api /api

EXPOSE 3000

CMD ["/api"]
