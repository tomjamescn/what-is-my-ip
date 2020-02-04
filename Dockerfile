FROM golang:alpine AS builder
WORKDIR /go/src/github.com/tomjamescn/what-is-my-ip
COPY main.go .
RUN go build -o main

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/tomjamescn/what-is-my-ip/main .

CMD ["./main"]

