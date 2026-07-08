FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ .
RUN CGO_ENABLED=0 go build -o /api ./cmd/

FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /api /api

EXPOSE 3000

CMD ["/bin/sh","-c","/api > /tmp/startup.log 2>&1; echo '---CRASH---' >> /tmp/startup.log; cat /tmp/startup.log; sleep 300"]
