FROM golang:1.17

RUN go env -w GOPROXY=https://goproxy.cn
RUN go get github.com/cespare/reflex

COPY reflex.conf /
COPY start.sh /

ENTRYPOINT [ "reflex", "-c", "/reflex.conf" ]