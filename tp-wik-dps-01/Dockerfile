FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /usr/local/bin/ping-server .

ENV PING_LISTEN_PORT=8080
EXPOSE 8080

USER appuser

ENTRYPOINT ["/usr/local/bin/ping-server"]
