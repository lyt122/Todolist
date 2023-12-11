FROM golang:1.21 AS builder

ENV GO111MODULE=on
ENV    CGO_ENABLED=0
ENV    GOOS=linux
ENV    GOARCH=amd64
ENV TZ  Asia/Shanghai
ENV GOPROXY  https://goproxy.cn,direct

RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
RUN go build -o output/main.exe main.go router.go router_gen.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
ENV service todolist
WORKDIR /app
COPY --from=builder /app/output /app/output
COPY --from=builder /app/config /app/config

CMD ["/app/output/main.exe"]
