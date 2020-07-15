# Build
FROM golang:alpine as builder

WORKDIR /bin
COPY ./server .
COPY .env .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o run

# Main
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /bin .

ENTRYPOINT ["./run"]