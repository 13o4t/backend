FROM golang:1.17 AS builder

COPY . /app

WORKDIR /app/cmd/myapp

RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download

RUN go build -trimpath -ldflags "-s -w" -o server .

FROM debian:buster-slim

COPY --from=builder /app/cmd/myapp/server /app/server

CMD ["/app/server"]