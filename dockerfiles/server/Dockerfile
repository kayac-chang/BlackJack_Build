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

# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

CMD ["./run"]